package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoPictureUpdate
修改图片名字
taobao.picture.update

修改指定图片的图片名 */
func TaobaoPictureUpdate(clt *core.SDKClient, req *media.TaobaoPictureUpdateAPIRequest, session string) (*media.TaobaoPictureUpdateAPIResponse, error) {
	var resp media.TaobaoPictureUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
