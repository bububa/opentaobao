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

// New
