package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingImagetranslateResultAPIRequest 获取图片翻译任务结果 API请求
// alibaba.seaking.imagetranslate.result
//
// 获取图片翻译任务结果
type AlibabaSeakingImagetranslateResultAPIRequest struct {
	model.Params
	// token来源站点
	_tokenFrom string
	// 任务id
	_taskId int64
	// 用户token
	_token string
}

// NewAlibabaSeakingImagetranslateResultRequest 初始化AlibabaSeakingImagetranslateResultAPIRequest对象
func NewAlibabaSeakingImagetranslateResultRequest() *AlibabaSeakingImagetranslateResultAPIRequest {
	return &AlibabaSeakingImagetranslateResultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagetranslateResultAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.imagetranslate.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagetranslateResultAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingImagetranslateResultAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// Get TokenFrom Getter
func (r AlibabaSeakingImagetranslateResultAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

// Set is TaskId Setter
// 任务id
func (r *AlibabaSeakingImagetranslateResultAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// Get TaskId Getter
func (r AlibabaSeakingImagetranslateResultAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// Set is Token Setter
// 用户token
func (r *AlibabaSeakingImagetranslateResultAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaSeakingImagetranslateResultAPIRequest) GetToken() string {
	return r._token
}
