package wlbimports

// WayBillDto 结构体
type WayBillDto struct {
	// 云打印数据
	CloudPrintData string `json:"cloud_print_data,omitempty" xml:"cloud_print_data,omitempty"`
	// 云打印pdf
	PdfWayBillUrl string `json:"pdf_way_bill_url,omitempty" xml:"pdf_way_bill_url,omitempty"`
}
