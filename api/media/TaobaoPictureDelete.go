package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopicturedelete 删除图片空间图片
// taobao.picture.delete
//
// 删除图片空间图片
func Taobaopicturedelete(clt *core.SDKClient, req *media.TaobaopicturedeleteAPIRequest, session string) (*media.TaobaopicturedeleteAPIResponse, error) {
	var resp media.TaobaopicturedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
