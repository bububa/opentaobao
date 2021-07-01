package aesolution

// OrderDetailQuery 结构体
type OrderDetailQuery struct {
	// It defines which details to be returned. Convert the number into binary format, for example, 16 to 10000. Only the last 5 bits take effects, starting from the end, 1st bit is for issue information, 2nd bit is for loan information, 3rd bit is for logistics information, 4th bit is for buyer information and 5th bit is for refund information. If any bit is 1, it means to return the corresponding information, for example, 3 wich is 00011, means to return issue information and logistics information. Leaving this field blank means return all information.
	ExtInfoBitFlag int64 `json:"ext_info_bit_flag,omitempty" xml:"ext_info_bit_flag,omitempty"`
	// Order id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
