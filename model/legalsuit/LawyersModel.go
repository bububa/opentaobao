package legalsuit

// LawyersModel 
type LawyersModel struct {
    // 律师列表
    Lawyers   []Lawyers `json:"lawyers,omitempty" xml:"lawyers>lawyers,omitempty"`
    // 操作类型
    OperationType   string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
}
