/*
Copyright 2022 QuanxiangCloud Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package search

import (
	"fmt"

	"github.com/quanxiang-cloud/initdb/pkg/toolkits"
)

type SearchConf struct {
	Host string `envconfig:"ESHOST"`
}

func (s *SearchConf) InitSearch(index string) error {
	userIndex := `
	{
        "settings": {
        "index": {
          "number_of_shards": 3,  
          "number_of_replicas": 2 
        }
      }
    }`
	URL := fmt.Sprintf("%s/%s?pretty", s.Host, index)
	return toolkits.DoPUT(URL, userIndex)
}
