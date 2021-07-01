package baichuan
import (
    "github.com/bububa/opentaobao/model"
)

// IsvItemSubDo 
type IsvItemSubDo struct {
    // 商品状态
    ItemStatus   *model.File `json:"item_status,omitempty" xml:"item_status,omitempty"`
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 添加时间
    AddTime   string `json:"add_time,omitempty" xml:"add_time,omitempty"`
    // 数据库索引id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
