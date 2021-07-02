package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPentaprismTaskQueryitemAPIRequest 查询任务当前进度 API请求
// taobao.pentaprism.task.queryitem
//
// 外网用户查询五棱镜任务系统当前进度
type TaobaoPentaprismTaskQueryitemAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// NewTaobaoPentaprismTaskQueryitemRequest 初始化TaobaoPentaprismTaskQueryitemAPIRequest对象
func NewTaobaoPentaprismTaskQueryitemRequest() *TaobaoPentaprismTaskQueryitemAPIRequest {
	return &TaobaoPentaprismTaskQueryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPentaprismTaskQueryitemAPIRequest) GetApiMethodName() string {
	return "taobao.pentaprism.task.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPentaprismTaskQueryitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOpenPo is OpenPo Setter
// TOP接口标准入参
func (r *TaobaoPentaprismTaskQueryitemAPIRequest) SetOpenPo(_openPo *OpenTaskPo) error {
	r._openPo = _openPo
	r.Set("open_po", _openPo)
	return nil
}

// GetOpenPo OpenPo Getter
func (r TaobaoPentaprismTaskQueryitemAPIRequest) GetOpenPo() *OpenTaskPo {
	return r._openPo
}
