package util

// QimenEvent 
/* model for simplify = false
type QimenEvent struct {

    // 奇门事件对象
    
    Event  *struct {
        Event  *Event `json:"event,omitempty"`
    } `json:"event,omitempty"`
    

}
*/

// QimenEvent 
type QimenEvent struct {

    // 奇门事件对象
    Event   *Event `json:"event,omitempty"`

}
