package portals

import "os"

const (
	api_post_url      = "api_post_url"
	api_image_url     = "api_image_url"
	api_post_user     = "api_post_user"
	api_post_password = "api_post_password"
)

var (
	api_url      = os.Getenv(api_post_url)
	api_img_url  = os.Getenv(api_image_url)
	api_user     = os.Getenv(api_post_user)
	api_password = os.Getenv(api_post_password)
	userAgent    = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/80.0.3987.149 Chrome/80.0.3987.149 Safari/537.36"
)

type ImgId struct {
	ID int64 `json:"id"`
}

type ApiField struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	Categories int8   `json:"categories"`
	Format     string `json:"format"`
	Feature    int64  `json:"featured_media"`
}

type Article struct {
	ArticleTitle   string
	URL            string
	ArticleContent string
	Category       int8
	ArticleImage   string
}
