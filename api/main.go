package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

type relatedProducts struct {
	Metadata struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Product struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"product"`
		ShowSizeFilter       bool `json:"showSizeFilter"`
		ShowColorFilter      bool `json:"showColorFilter"`
		ShowHeelHeightFilter bool `json:"showHeelHeightFilter"`
		ShowConditionFilter  bool `json:"showConditionFilter"`
	} `json:"metadata"`
	Products []struct {
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		BrandedName    string  `json:"brandedName"`
		UnbrandedName  string  `json:"unbrandedName"`
		Currency       string  `json:"currency"`
		Price          float64 `json:"price"`
		PriceLabel     string  `json:"priceLabel"`
		SalePrice      float64 `json:"salePrice"`
		SalePriceLabel string  `json:"salePriceLabel"`
		InStock        bool    `json:"inStock"`
		Stock          []struct {
			Color struct {
				Name string `json:"name"`
			} `json:"color"`
			Size struct {
				Name string `json:"name"`
			} `json:"size"`
		} `json:"stock"`
		Retailer struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Logo struct {
				Sizes struct {
					Mobile struct {
						URL string `json:"url"`
					} `json:"mobile"`
					Default struct {
						URL string `json:"url"`
					} `json:"default"`
					Featured struct {
						URL string `json:"url"`
					} `json:"featured"`
					Mobile2X struct {
						URL string `json:"url"`
					} `json:"mobile@2x"`
					Featured2X struct {
						URL string `json:"url"`
					} `json:"featured@2x"`
					Default2X struct {
						URL string `json:"url"`
					} `json:"default@2x"`
				} `json:"sizes"`
			} `json:"logo"`
			UserID                 int    `json:"userId"`
			UserHandle             string `json:"userHandle"`
			UserImage              string `json:"userImage"`
			URLIdentifier          string `json:"urlIdentifier"`
			FavIcon                string `json:"favIcon"`
			EffectiveContentLocale string `json:"effectiveContentLocale"`
			Handle                 string `json:"handle"`
		} `json:"retailer"`
		Brand struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Logo struct {
				Sizes struct {
					Mobile struct {
						URL string `json:"url"`
					} `json:"mobile"`
					Default struct {
						URL string `json:"url"`
					} `json:"default"`
					Featured struct {
						URL string `json:"url"`
					} `json:"featured"`
					Mobile2X struct {
						URL string `json:"url"`
					} `json:"mobile@2x"`
					Featured2X struct {
						URL string `json:"url"`
					} `json:"featured@2x"`
					Default2X struct {
						URL string `json:"url"`
					} `json:"default@2x"`
				} `json:"sizes"`
			} `json:"logo"`
			UserID              int    `json:"userId"`
			UserHandle          string `json:"userHandle"`
			UserImage           string `json:"userImage"`
			URLIdentifier       string `json:"urlIdentifier"`
			LegacyURLIdentifier string `json:"legacyUrlIdentifier"`
		} `json:"brand"`
		Locale        string `json:"locale"`
		Description   string `json:"description"`
		ClickURL      string `json:"clickUrl"`
		PageURL       string `json:"pageUrl"`
		ClusterURL    string `json:"clusterUrl"`
		DirectURL     string `json:"directUrl"`
		URLIdentifier string `json:"urlIdentifier"`
		Image         struct {
			Sizes struct {
				Small struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Small"`
				XLarge struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"XLarge"`
				Medium struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Medium"`
				Large struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Large"`
				IPhoneSmall struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"IPhoneSmall"`
				Best struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Best"`
				Original struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Original"`
				IPhone struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"IPhone"`
			} `json:"sizes"`
			ID string `json:"id"`
		} `json:"image"`
		AlternateImages []struct {
			Sizes struct {
				Small struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Small"`
				XLarge struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"XLarge"`
				Medium struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Medium"`
				Large struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Large"`
				IPhoneSmall struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"IPhoneSmall"`
				Best struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Best"`
				Original struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"Original"`
				IPhone struct {
					SizeName     string `json:"sizeName"`
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					ActualWidth  int    `json:"actualWidth"`
					ActualHeight int    `json:"actualHeight"`
				} `json:"IPhone"`
			} `json:"sizes"`
			ID string `json:"id"`
		} `json:"alternateImages"`
		CheckoutEnabled bool `json:"checkoutEnabled"`
		Colors          []struct {
			Name  string `json:"name"`
			Image struct {
				Sizes struct {
					Small struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"Small"`
					XLarge struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"XLarge"`
					Medium struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"Medium"`
					Large struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"Large"`
					IPhoneSmall struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"IPhoneSmall"`
					Best struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"Best"`
					Original struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"Original"`
					IPhone struct {
						SizeName     string `json:"sizeName"`
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						ActualWidth  int    `json:"actualWidth"`
						ActualHeight int    `json:"actualHeight"`
					} `json:"IPhone"`
				} `json:"sizes"`
				ID string `json:"id"`
			} `json:"image"`
			SwatchURL       string `json:"swatchUrl"`
			CanonicalColors []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Rank int    `json:"rank"`
			} `json:"canonicalColors"`
		} `json:"colors"`
		Sizes []struct {
			Name          string `json:"name"`
			CanonicalSize struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"canonicalSize"`
		} `json:"sizes"`
		Categories []struct {
			NumID       int    `json:"numId"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			ShortName   string `json:"shortName"`
			FullName    string `json:"fullName"`
			LocalizedID string `json:"localizedId"`
		} `json:"categories"`
		ExtractDate     string        `json:"extractDate"`
		Badges          []interface{} `json:"badges"`
		PromotionalDeal struct {
			ID           int    `json:"id"`
			Type         string `json:"type"`
			TypeLabel    string `json:"typeLabel"`
			Title        string `json:"title"`
			ShortTitle   string `json:"shortTitle"`
			CheckoutCode string `json:"checkoutCode"`
			Retailer     struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Logo struct {
					Sizes struct {
						Mobile struct {
							URL string `json:"url"`
						} `json:"mobile"`
						Default struct {
							URL string `json:"url"`
						} `json:"default"`
						Featured struct {
							URL string `json:"url"`
						} `json:"featured"`
						Mobile2X struct {
							URL string `json:"url"`
						} `json:"mobile@2x"`
						Featured2X struct {
							URL string `json:"url"`
						} `json:"featured@2x"`
						Default2X struct {
							URL string `json:"url"`
						} `json:"default@2x"`
					} `json:"sizes"`
				} `json:"logo"`
				UserID                 int    `json:"userId"`
				UserHandle             string `json:"userHandle"`
				UserImage              string `json:"userImage"`
				URLIdentifier          string `json:"urlIdentifier"`
				FavIcon                string `json:"favIcon"`
				EffectiveContentLocale string `json:"effectiveContentLocale"`
				Handle                 string `json:"handle"`
				DeeplinkSupport        bool   `json:"deeplinkSupport"`
				HostDomain             string `json:"hostDomain"`
			} `json:"retailer"`
			StartDate struct {
				Date      string `json:"date"`
				Timestamp int    `json:"timestamp"`
				Friendly  string `json:"friendly"`
			} `json:"startDate"`
			EndDate struct {
				Date      string `json:"date"`
				Timestamp int    `json:"timestamp"`
				Friendly  string `json:"friendly"`
			} `json:"endDate"`
		} `json:"promotionalDeal"`
		FavoriteCount int    `json:"favoriteCount"`
		HasFavorite   bool   `json:"hasFavorite"`
		SeeMoreLabel  string `json:"seeMoreLabel"`
		SeeMoreURL    string `json:"seeMoreUrl"`
		PreOwned      bool   `json:"preOwned"`
		Rental        bool   `json:"rental"`
	} `json:"products"`
}

func main() {
	relatedProductsMap := parseRelatedProducts()

	server := gin.Default()
	server.GET("/related/:name", func(c *gin.Context) {
		id := c.Param("name")
		productList, productListPresent := relatedProductsMap[id]

		if productListPresent {
			client := &http.Client{}
			request, error := http.NewRequest("GET", "https://api.shopstyle.com/api/v2/products", nil)
			if error != nil {
				panic(error)
			}

			query := request.URL.Query()
			query.Add("pid", "shopstyle")

			for _, productID := range productList {
				query.Add("prodid", productID)
			}

			request.URL.RawQuery = query.Encode()
			response, err := client.Do(request)

			if err != nil {
				panic(err)
			}

			defer response.Body.Close()

			var body relatedProducts
			json.NewDecoder(response.Body).Decode(&body)
			c.JSON(http.StatusOK, body)
		} else {
			c.JSON(http.StatusOK, gin.H{"products": "[]"})
		}

	})

	server.Run(":8080")
}

func parseRelatedProducts() map[string][]string {
	file, err := os.Open("related-products-dictionary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	relatedProducts := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "|")
		prodIDKey := splitLine[0]
		prodIds := strings.Split(splitLine[1], ";")
		relatedProducts[prodIDKey] = prodIds
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return relatedProducts
}
