package icbudropshipping

// LadderPeriod 
type LadderPeriod struct {

    // max quantity
    
    MaxQuantity   int64 `json:"max_quantity,omitempty" xml:"max_quantity,omitempty"`
    

    // min quantity
    
    MinQuantity   int64 `json:"min_quantity,omitempty" xml:"min_quantity,omitempty"`
    

    // Delivery time
    
    ProcessPeriod   int64 `json:"process_period,omitempty" xml:"process_period,omitempty"`
    

}
