package logistic

// ChainStore 
type ChainStore struct {

    // 门店code
    ChainstoreCode   string `json:"chainstore_code,omitempty"`

    // 经度
    Longitude   string `json:"longitude,omitempty"`

    // 纬度
    Latitude   string `json:"latitude,omitempty"`

}
