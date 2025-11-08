package controllers
import(
	"Invoice-Payment-System/internal/domain"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InvoiceHandler struct{
	InvoiceUsecase domain.InvoiceUsecase
	Validate *validator.Validate
}

func NewInvoiceController(invoiceUsecase domain.InvoiceUsecase) *InvoiceHandler{
	return &InvoiceHandler{
		InvoiceUsecase: invoiceUsecase,
		Validate: validator.New(),
	}
}

func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	var invoice domain.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.Validate.Struct(invoice); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	
	paymentURL, err := h.InvoiceUsecase.CreateInvoice(&invoice)
	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":      "Invoice created successfully",
		"invoice":      invoice,
		"payment_link": paymentURL,
	})
}
func (h *InvoiceHandler) GetInvoiceByID(c *gin.Context) {
	id := c.Param("id")
	invoice, err := h.InvoiceUsecase.GetInvoice(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"invoice": invoice})
}

func (h *InvoiceHandler) GetInvoicesBySender(c *gin.Context) {
	email := c.Query("email")
	
	invoices, err := h.InvoiceUsecase.GetInvoicesBySender(email)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, invoices)
}
