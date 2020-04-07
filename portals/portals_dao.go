package portals

import (
	"fmt"

	"github.com/amisini/scraping_portals/date_utils"
	"github.com/amisini/scraping_portals/portals_db"
)

const (
	queryInsertUser = (`INSERT INTO tbl_artikuj(
      Source_Id,
      Kategori_Id,
      Titulli,
      Permbajtja,
      Url,
      ImageUrl,
      DateInserted,
      Active)
      VALUES (?,?,?,?,?,?,?,?);`)
)

func (article *Article) Save() error {
	stmt, err := portals_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(1,
		article.Category,
		article.ArticleTitle,
		article.ArticleContent,
		article.URL,
		article.ArticleImage,
		date_utils.GetNowString(), 1)

	if err != nil {
		return err
	}

	articleId, err := insertResult.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(articleId)

	return nil
}
