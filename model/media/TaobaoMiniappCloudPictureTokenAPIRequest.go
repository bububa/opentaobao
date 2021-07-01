package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudPictureTokenAPIRequest
云存储上传token获取 API请求
taobao.miniapp.cloud.picture.token

获取图片上传token，作为http://upload.cloud.tmall.com/api/proxy/upload接口上传的凭证。

上传的html示例代码为：
&lt;form action=&quot;http://upload.cloud.tmall.com/api/proxy/upload&quot; method=&quot;post&quot; enctype=&quot;multipart/form-data&quot;&gt;
	上传token: &lt;input type=&quot;text&quot; name=&quot;Authorization&quot; value=&quot;&quot; /&gt;&lt;br/&gt;
	上传文件:&lt;input type=&quot;file&quot; name=&quot;content&quot; /&gt;&lt;br/&gt;
	&lt;input type=&quot;submit&quot; value=&quot;Submit &quot;/&gt;
&lt;/form&gt; */
type TaobaoMiniappCloudPictureTokenAPIRequest struct {
	model.Params
	// 请求参数
	_generateTokenRequest *GenerateTokenRequest
}

// NewTaobaoMiniappCloudPictureTokenRequest 初始化TaobaoMiniappCloudPictureTokenAPIRequest对象
func NewTaobaoMiniappCloudPictureTokenRequest() *TaobaoMiniappCloudPictureTokenAPIRequest {
	return &TaobaoMiniappCloudPictureTokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.picture.token"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GenerateTokenRequest Setter
// 请求参数
func (r *TaobaoMiniappCloudPictureTokenAPIRequest) SetGenerateTokenRequest(_generateTokenRequest *GenerateTokenRequest) error {
	r._generateTokenRequest = _generateTokenRequest
	r.Set("generate_token_request", _generateTokenRequest)
	return nil
}

// Get GenerateTokenRequest Getter
func (r TaobaoMiniappCloudPictureTokenAPIRequest) GetGenerateTokenRequest() *GenerateTokenRequest {
	return r._generateTokenRequest
}
