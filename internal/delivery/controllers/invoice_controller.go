package controller
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
	if err := c.ShouldBindBodyWithJSON(&invoice); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.Validate.Struct(invoice); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	
	err := h.InvoiceUsecase.CreateInvoice(&invoice)
	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, invoice)
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
