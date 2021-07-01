package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoMiniappCloudPictureToken
云存储上传token获取
taobao.miniapp.cloud.picture.token

获取图片上传token，作为http://upload.cloud.tmall.com/api/proxy/upload接口上传的凭证。

上传的html示例代码为：
&lt;form action=&quot;http://upload.cloud.tmall.com/api/proxy/upload&quot; method=&quot;post&quot; enctype=&quot;multipart/form-data&quot;&gt;
	上传token: &lt;input type=&quot;text&quot; name=&quot;Authorization&quot; value=&quot;&quot; /&gt;&lt;br/&gt;
	上传文件:&lt;input type=&quot;file&quot; name=&quot;content&quot; /&gt;&lt;br/&gt;
	&lt;input type=&quot;submit&quot; value=&quot;Submit &quot;/&gt;
&lt;/form&gt; */
func TaobaoMiniappCloudPictureToken(clt *core.SDKClient, req *media.TaobaoMiniappCloudPictureTokenAPIRequest, session string) (*media.TaobaoMiniappCloudPictureTokenAPIResponse, error) {
	var resp media.TaobaoMiniappCloudPictureTokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
