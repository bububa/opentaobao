package btrip

// BtripHotelCancelOrderRs 
type BtripHotelCancelOrderRs struct {

    // 是否取消成功
    
    CancelSuccess   bool `json:"cancel_success,omitempty" xml:"cancel_success,omitempty"`
    

    // 罚金
    
    ForfeitFee   int64 `json:"forfeit_fee,omitempty" xml:"forfeit_fee,omitempty"`
    

}
