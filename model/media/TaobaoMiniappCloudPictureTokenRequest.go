package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储上传token获取 API请求
taobao.miniapp.cloud.picture.token

获取图片上传token，作为http://upload.cloud.tmall.com/api/proxy/upload接口上传的凭证。

上传的html示例代码为：
&lt;form action=&quot;http://upload.cloud.tmall.com/api/proxy/upload&quot; method=&quot;post&quot; enctype=&quot;multipart/form-data&quot;&gt;
	上传token: &lt;input type=&quot;text&quot; name=&quot;Authorization&quot; value=&quot;&quot; /&gt;&lt;br/&gt;
	上传文件:&lt;input type=&quot;file&quot; name=&quot;content&quot; /&gt;&lt;br/&gt;
	&lt;input type=&quot;submit&quot; value=&quot;Submit &quot;/&gt;
&lt;/form&gt;
*/
type TaobaoMiniappCloudPictureTokenRequest struct {
    model.Params
    // 请求参数
    _generateTokenRequest   *GenerateTokenRequest
}

// 初始化TaobaoMiniappCloudPictureTokenRequest对象
func NewTaobaoMiniappCloudPictureTokenRequest() *TaobaoMiniappCloudPictureTokenRequest{
    return &TaobaoMiniappCloudPictureTokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudPictureTokenRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.picture.token"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudPictureTokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GenerateTokenRequest Setter
// 请求参数
func (r *TaobaoMiniappCloudPictureTokenRequest) SetGenerateTokenRequest(_generateTokenRequest *GenerateTokenRequest) error {
    r._generateTokenRequest = _generateTokenRequest
    r.Set("generate_token_request", _generateTokenRequest)
    return nil
}

// GenerateTokenRequest Getter
func (r TaobaoMiniappCloudPictureTokenRequest) GetGenerateTokenRequest() *GenerateTokenRequest {
    return r._generateTokenRequest
}
