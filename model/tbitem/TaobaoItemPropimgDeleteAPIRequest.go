package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPropimgDeleteAPIRequest 删除属性图片 API请求
// taobao.item.propimg.delete
//
// 删除propimg_id 所指定的商品属性图片 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;propimg_id对应的属性图片需要属于num_iid对应的商品
type TaobaoItemPropimgDeleteAPIRequest struct {
	model.Params
	// 商品数字ID，必选
	_numIid int64
	// 商品属性图片ID
	_id int64
}

// NewTaobaoItemPropimgDeleteRequest 初始化TaobaoItemPropimgDeleteAPIRequest对象
func NewTaobaoItemPropimgDeleteRequest() *TaobaoItemPropimgDeleteAPIRequest {
	return &TaobaoItemPropimgDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemPropimgDeleteAPIRequest) Reset() {
	r._numIid = 0
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemPropimgDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.propimg.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemPropimgDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemPropimgDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemPropimgDeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemPropimgDeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetId is Id Setter
// 商品属性图片ID
func (r *TaobaoItemPropimgDeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoItemPropimgDeleteAPIRequest) GetId() int64 {
	return r._id
}

var poolTaobaoItemPropimgDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemPropimgDeleteRequest()
	},
}

// GetTaobaoItemPropimgDeleteRequest 从 sync.Pool 获取 TaobaoItemPropimgDeleteAPIRequest
func GetTaobaoItemPropimgDeleteAPIRequest() *TaobaoItemPropimgDeleteAPIRequest {
	return poolTaobaoItemPropimgDeleteAPIRequest.Get().(*TaobaoItemPropimgDeleteAPIRequest)
}

// ReleaseTaobaoItemPropimgDeleteAPIRequest 将 TaobaoItemPropimgDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemPropimgDeleteAPIRequest(v *TaobaoItemPropimgDeleteAPIRequest) {
	v.Reset()
	poolTaobaoItemPropimgDeleteAPIRequest.Put(v)
}
