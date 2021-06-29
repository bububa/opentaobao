package xhotelitem
import (
    "github.com/bububa/opentaobao/model"
)

// RoomType 
type RoomType struct {
    // 房型名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 阿里房型id
    Rid   int64 `json:"rid,omitempty" xml:"rid,omitempty"`
    // 房型状态。0:正常，-1:删除，-2:停售
    Status   *model.File `json:"status,omitempty" xml:"status,omitempty"`
    // 系统商，一般不填写，使用须申请
    Vendor   string `json:"vendor,omitempty" xml:"vendor,omitempty"`
    // 商家房型ID
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 房价列表
    RatePlanList   []RatepPlan `json:"rate_plan_list,omitempty" xml:"rate_plan_list>ratep_plan,omitempty"`
}
