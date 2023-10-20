package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest 信息流单元内解绑资源位 API请求
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
type TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest struct {
	model.Params
	// 广告位id
	_adzoneIdList []string
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemAdgroupAdzoneUnbindRequest 初始化TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest对象
func NewTaobaoFeedflowItemAdgroupAdzoneUnbindRequest() *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest {
	return &TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) Reset() {
	r._adzoneIdList = r._adzoneIdList[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneIdList is AdzoneIdList Setter
// 广告位id
func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) SetAdzoneIdList(_adzoneIdList []string) error {
	r._adzoneIdList = _adzoneIdList
	r.Set("adzone_id_list", _adzoneIdList)
	return nil
}

// GetAdzoneIdList AdzoneIdList Getter
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) GetAdzoneIdList() []string {
	return r._adzoneIdList
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupAdzoneUnbindRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupAdzoneUnbindRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest
func GetTaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest() *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest {
	return poolTaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest.Get().(*TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest 将 TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest(v *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest.Put(v)
}
