package logistic

// ExtensionTopDto 结构体
type ExtensionTopDto struct {
	// Whether a seller has added invoice to transaction order. If a seller doesn't provide invoice, Correios will be the only available shipment option. This is mainly to remind sellers of adding invoice before shipment
	InvoiceAdded bool `json:"invoice_added,omitempty" xml:"invoice_added,omitempty"`
}
