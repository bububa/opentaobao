package qimen

// ResponseDO 
/* model for simplify = false
type ResponseDO struct {

    // 响应结果:success|failure,success,string(10),必填,
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码,0,string(50),,
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息,invalid appkey,string(100),,
    
    Message   string `json:"message,omitempty"`
    

    // itemInventories
    
    ItemInventories  struct {
        ItemInventory  []ItemInventory `json:"item_inventory,omitempty"`
    } `json:"itemInventories,omitempty"`
    

    // item
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

// ResponseDO 
type ResponseDO struct {

    // 响应结果:success|failure,success,string(10),必填,
    Flag   string `json:"flag,omitempty"`

    // 响应码,0,string(50),,
    Code   string `json:"code,omitempty"`

    // 响应信息,invalid appkey,string(100),,
    Message   string `json:"message,omitempty"`

    // itemInventories
    ItemInventories   []ItemInventory `json:"itemInventories,omitempty"`

    // item
    Item   *Item `json:"item,omitempty"`

}
