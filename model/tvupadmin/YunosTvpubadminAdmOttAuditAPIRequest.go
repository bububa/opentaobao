package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminAdmOttAuditAPIRequest
优酷OTT广告素材审核 API请求
yunos.tvpubadmin.adm.ott.audit

审核优酷OTT端广告素材 */
type YunosTvpubadminAdmOttAuditAPIRequest struct {
	model.Params
	// 广告审核内容，json格式
	_data string
}

// NewYunosTvpubadminAdmOttAuditRequest 初始化YunosTvpubadminAdmOttAuditAPIRequest对象
func NewYunosTvpubadminAdmOttAuditRequest() *YunosTvpubadminAdmOttAuditAPIRequest {
	return &YunosTvpubadminAdmOttAuditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.adm.ott.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Data Setter
// 广告审核内容，json格式
func (r *YunosTvpubadminAdmOttAuditAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetData() string {
	return r._data
}
