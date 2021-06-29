package aesolution

// SynchronizeProductResponseDTO 
type SynchronizeProductResponseDTO struct {
    // error code
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // error message
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // product id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
