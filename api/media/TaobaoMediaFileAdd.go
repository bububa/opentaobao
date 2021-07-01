package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoMediaFileAdd
多媒体平台文件添加
taobao.media.file.add

用户通过top上传文件到多媒体平台 */
func TaobaoMediaFileAdd(clt *core.SDKClient, req *media.TaobaoMediaFileAddAPIRequest, session string) (*media.TaobaoMediaFileAddAPIResponse, error) {
	var resp media.TaobaoMediaFileAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
