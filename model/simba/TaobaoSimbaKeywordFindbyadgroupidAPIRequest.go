package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordfindbyadgroupidAPIRequest 根据推广单元id获取关键词 API请求
// taobao.simba.keyword.findbyadgroupid
//
// 根据一个关键词Id列表取得一组关键词
type TaobaosimbakeywordfindbyadgroupidAPIRequest struct {
	model.Params
	// 推广单元id
	_adgroupId int64
}

// NewTaobaosimbakeywordfindbyadgroupidRequest 初始化TaobaosimbakeywordfindbyadgroupidAPIRequest对象
func NewTaobaosimbakeywordfindbyadgroupidRequest() *TaobaosimbakeywordfindbyadgroupidAPIRequest {
	return &TaobaosimbakeywordfindbyadgroupidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordfindbyadgroupidAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.findbyadgroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordfindbyadgroupidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordfindbyadgroupidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbakeywordfindbyadgroupidAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordfindbyadgroupidAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
