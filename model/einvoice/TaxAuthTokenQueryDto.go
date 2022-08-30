package einvoice

// TaxAuthTokenQueryDto 结构体
type TaxAuthTokenQueryDto struct {
	// 税控产品码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 服务商平台授权商户的税号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
}
