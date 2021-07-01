package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItempropvaluesGetAPIRequest
获取标准类目属性值 API请求
taobao.itempropvalues.get

获取标准类目属性值 */
type TaobaoItempropvaluesGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
	_fields []string
	// 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
	_cid int64
	// 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
	_pvs string
	// 获取类目的类型：1代表集市、2代表天猫
	_type int64
	// 属性的Key，支持多条，以“,”分隔
	_attrKeys []string
}

// New
