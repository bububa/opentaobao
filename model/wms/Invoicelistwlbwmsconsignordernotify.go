package wms

// Invoicelistwlbwmsconsignordernotify 
/* model for simplify = false
type Invoicelistwlbwmsconsignordernotify struct {

    // 发票信息
    
    InvoiceInfo  *struct {
        Invoicewlbwmsconsignordernotify  *Invoicewlbwmsconsignordernotify `json:"invoicewlbwmsconsignordernotify,omitempty"`
    } `json:"invoice_info,omitempty"`
    

}
*/

// Invoicelistwlbwmsconsignordernotify 
type Invoicelistwlbwmsconsignordernotify struct {

    // 发票信息
    InvoiceInfo   *Invoicewlbwmsconsignordernotify `json:"invoice_info,omitempty"`

}
