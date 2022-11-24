forEach: RelationCommandInfo
fileName: {{commandValue.aggregate.namePascalCase}}Service.go
path: {{boundedContext.name}}/external
---
package external 

import (
	"github.com/go-resty/resty/v2"
	"fmt"
	"{{boundedContext.name}}/config"
)

var client = resty.New()

{{#if commandValue.restRepositoryInfo.getMethod}}
func Get{{commandValue.aggregate.namePascalCase}}( {{commandValue.aggregate.keyFieldDescriptor.name}} {{#typeCheck commandValue.aggregate.keyFieldDescriptor.className}} {{/typeCheck}}) *resty.Response {
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" , {{commandValue.aggregate.keyFieldDescriptor.name}} )
	resp, _ := client.R().Get(target)

	return resp
}
{{else if commandValue.restRepositoryInfo.method}}
{{#commandValue.isRestRepository}}
func {{commandValue.nameCamelCase}}({{commandValue.aggregate.nameCamelCase}} {{commandValue.aggregate.namePascalCase}}) *resty.Response {
	{{#MethodPost commandValue.restRepositoryInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" )
	resp, _ := client.R().SetBody({{commandValue.aggregate.nameCamelCase}}).Post(target)

	return resp
	{{/MethodPost}}

	{{#MethodGet commandValue.restRepositoryInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" , {{commandValue.aggregate.nameCamelCase}} )
	resp, _ := client.R().Get(target)
	return resp
	{{/MethodGet}}

	{{#MethodUpdate commandValue.restRepositoryInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" )
	resp, _ := client.R().SetBody({{commandValue.aggregate.nameCamelCase}}).Put(target)
	return resp
	{{/MethodUpdate}}

	{{#MethodDelete commandValue.restRepositoryInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" , {{commandValue.aggregate.nameCamelCase}} )
	resp, _ := client.R().Delete(target)

	return resp
	{{/MethodDelete}}
}
{{/commandValue.isRestRepository}}
{{^commandValue.isRestRepository}}
func {{commandValue.namePascalCase}}({{commandValue.aggregate.nameCamelCase}} {{commandValue.aggregate.namePascalCase}}) *resty.Response{
	{{#MethodPost commandValue.controllerInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" )
	resp, _ := client.R().SetBody({{commandValue.aggregate.nameCamelCase}}).Post(target)

	return resp
	{{/MethodPost}}
	{{#MethodGet commandValue.controllerInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" ,{{commandValue.aggregate.nameCamelCase}}.{{commandValue.aggregate.keyFieldDescriptor.namePascalCase}} ,{{commandValue.controllerInfo.apiPath}} )
	resp, _ := client.R().Get(target)

	return resp
	{{/MethodGet}}
	{{#MethodUpdate commandValue.controllerInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" )
	resp, _ := client.R().SetBody({{commandValue.aggregate.nameCamelCase}}).Put(target)

	return resp
	{{/MethodUpdate}}
	{{#MethodDelete commandValue.controllerInfo.method}}
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" , {{commandValue.aggregate.nameCamelCase}}.{{commandValue.aggregate.keyFieldDescriptor.namePascalCase}} ,{{commandValue.controllerInfo.apiPath}})
	resp, _ := client.R().Delete(target)

	return resp
	{{/MethodDelete}}
}
{{/commandValue.isRestRepository}}
{{else}}
func Get{{commandValue.aggregate.namePascalCase}}( {{commandValue.aggregate.keyFieldDescriptor.name}} {{#typeCheck commandValue.aggregate.keyFieldDescriptor.className}} {{/typeCheck}}) *resty.Response{
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s", options["api_url_{{commandValue.boundedContext.name}}"], "{{commandValue.aggregate.namePlural}}" , {{commandValue.aggregate.keyFieldDescriptor.name}} )
	resp, _ := client.R().Get(target)

	return resp
}
{{/if}}

<function>
	window.$HandleBars.registerHelper('MethodGet', function(method, options){
        if(method && method.endsWith('GET')){
        	return options.fn(this)
		}
		else{
			return options.inverse(this)
		}
    });
	window.$HandleBars.registerHelper('MethodPost', function(method, options){
        if(method && method.endsWith('POST')){
        	return options.fn(this)
		}
		else{
			return options.inverse(this)
		}
    });
	window.$HandleBars.registerHelper('MethodUpdate', function(method, options){
        if(method.endsWith('PUT')){
        	return options.fn(this)
		}
		else{
			return options.inverse(this)
		}
    });
	window.$HandleBars.registerHelper('MethodDelete', function(method, options){
        if(method.endsWith('DELETE')){
        	return options.fn(this)
		}
		else{
			return options.inverse(this)
		}
    });
	 window.$HandleBars.registerHelper('typeCheck', function (className) {
        if(className.endsWith("String")){
            return "string"
        }
		else if(className.endsWith("Integer")){
			return "int"
		}
		else if(className.endsWith("Float")){
			return "float64"
		}
		else if(className.endsWith("Long")){
			return "int"
		}
		else if(className.endsWith("Boolean")){
			return "bool"
		}
		else if(className.endsWith("Double")){
			return "int"
		}

    });
</function>