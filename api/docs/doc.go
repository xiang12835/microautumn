package docs

//bee run -gendoc=true -downdoc=true

import (
	"microautumn/api/docs/swagger"
)

type APIDoc struct {
	Api    swagger.ResourceListing
	Subapi map[string][]swagger.API
}

func init() {

}
