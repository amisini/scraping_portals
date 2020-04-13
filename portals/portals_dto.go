package portals

import (
	"os"
	"strings"
)

const (
	api_post_url                    = "api_post_url"
	api_image_url                   = "api_image_url"
	api_telegrafi_post_user         = "api_telegrafi_post_user"
	api_telegrafi_post_password     = "api_telegrafi_post_password"
	api_gazetaexpress_post_user     = "api_gazetaexpress_post_user"
	api_gazetaexpress_post_password = "api_gazetaexpress_post_password"
	api_indeksonline_post_user      = "api_indeksonline_post_user"
	api_indeksonline_post_password  = "api_indeksonline_post_password"
)

var (
	api_url               = os.Getenv(api_post_url)
	api_img_url           = os.Getenv(api_image_url)
	api_telegrafi_user    = os.Getenv(api_telegrafi_post_user)
	api_telgrafi_password = os.Getenv(api_telegrafi_post_password)

	api_gazetaexpress_user     = os.Getenv(api_gazetaexpress_post_user)
	api_gazetaexpress_password = os.Getenv(api_gazetaexpress_post_password)

	api_indeksonline_user     = os.Getenv(api_indeksonline_post_user)
	api_indeksonline_password = os.Getenv(api_indeksonline_post_password)

	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/80.0.3987.149 Chrome/80.0.3987.149 Safari/537.36"
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
	PortalID       int8
	ArticleTitle   string
	URL            string
	ArticleContent string
	Category       int8
	ArticleImage   string
}

var categories = map[string]int8{
	"lajme":       1,
	"sport":       2,
	"magazina":    3,
	"showbiz":     3,
	"muzik":       3,
	"roz":         3,
	"yjet":        3,
	"teknologji":  4,
	"tech":        4,
	"kuriozitete": 5,
	"shendetesi":  6,
	"shëndetësi":  6,
	"shneta":      6,
	"ekonomi":     7,
}

func GetCategory(categories map[string]int8, cat string) int8 {
	for key, value := range categories {
		if strings.Contains(strings.ToLower(cat), key) {
			return value
		}
	}
	return 0
}

func GetPortalCredentials(cred string) (string, string) {
	switch cred {
	case "telegrafi":
		return api_telegrafi_user, api_telgrafi_password
	case "gazetaexpress":
		return api_gazetaexpress_user, api_gazetaexpress_password
	case "indeksonline":
		return api_indeksonline_user, api_indeksonline_password
	default:
		return "", ""
	}
}
