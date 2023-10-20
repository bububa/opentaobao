package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardpagetmpAPIRequest 查询卡模板列表(支持数据下行) API请求
// alibaba.alsc.crm.card.pagetmp
//
// 查询卡模板列表(支持数据下行)
// 当传递了数据下行参数:
//   - isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
//   - 否则分页查询卡模板,返回结果带有分页信息
type AlibabaalsccrmcardpagetmpAPIRequest struct {
	model.Params
	// 请求结果
	_paramPullCardTemplateOpenReq *PullCardTemplateOpenReq
}

// NewAlibabaalsccrmcardpagetmpRequest 初始化AlibabaalsccrmcardpagetmpAPIRequest对象
func NewAlibabaalsccrmcardpagetmpRequest() *AlibabaalsccrmcardpagetmpAPIRequest {
	return &AlibabaalsccrmcardpagetmpAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardpagetmpAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.pagetmp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardpagetmpAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardpagetmpAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPullCardTemplateOpenReq is ParamPullCardTemplateOpenReq Setter
// 请求结果
func (r *AlibabaalsccrmcardpagetmpAPIRequest) SetParamPullCardTemplateOpenReq(_paramPullCardTemplateOpenReq *PullCardTemplateOpenReq) error {
	r._paramPullCardTemplateOpenReq = _paramPullCardTemplateOpenReq
	r.Set("param_pull_card_template_open_req", _paramPullCardTemplateOpenReq)
	return nil
}

// GetParamPullCardTemplateOpenReq ParamPullCardTemplateOpenReq Getter
func (r AlibabaalsccrmcardpagetmpAPIRequest) GetParamPullCardTemplateOpenReq() *PullCardTemplateOpenReq {
	return r._paramPullCardTemplateOpenReq
}
