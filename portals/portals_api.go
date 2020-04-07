package portals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/amisini/scraping_portals/file_utils"
)

func (article *Article) SaveAPI() error {
	fileUrl := article.ArticleImage
	segments := strings.Split(fileUrl, "/")
	fileName := segments[len(segments)-1]

	fileSave := "tmp/" + fileName

	if err := file_utils.DownloadFile(fileSave, fileUrl); err != nil {
		return err
	}

	reqImg, err := file_utils.NewfileUploadRequest(api_img_url, "file", fileSave)
	if err != nil {
		return err
	}

	reqImg.SetBasicAuth(api_user, api_password)
	reqImg.Header.Set("Content-Disposition", "attachment; filename="+fileName)
	reqImg.Header.Set("User-Agent", userAgent)

	clientImg := &http.Client{}
	respImg, err := clientImg.Do(reqImg)
	if err != nil {
		return err
	}
	defer respImg.Body.Close()
	body, err := ioutil.ReadAll(respImg.Body)
	if err != nil {
		return err
	}

	var imgId ImgId
	json.Unmarshal(body, &imgId)

	jsonStr, err := json.Marshal(ApiField{
		Title:      article.ArticleTitle,
		Content:    article.ArticleContent,
		Status:     "publish",
		Categories: article.Category,
		Format:     "standard",
		Feature:    imgId.ID})

	if err != nil {
		return err
	}

	delete := os.Remove(fileSave)
	if delete != nil {
		return delete
	}

	req, err := http.NewRequest("POST", api_url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.SetBasicAuth(api_user, api_password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status", resp.Status)
	return nil
}
