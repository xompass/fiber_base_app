package swagger_defaults

import "html/template"

var DefaultTheme = template.CSS(`html {
    box-sizing: border-box;
    overflow: -moz-scrollbars-vertical;
    overflow-y: scroll;
}

*,
*:before,
*:after {
    box-sizing: inherit;
}

body {
    margin: 0;
    background: #fafafa;
}

.swagger-ui .topbar {
    background: #2E363F;
    padding: 15px 0;
}

.swagger-ui .wrapper {
    max-width: 1460px;
}

.swagger-ui .topbar a {
    height: 40px;
    vertical-align: top;
    display: inline-block;
    background-image: url(https://admin-xompass.azureedge.net/assets/images/logos/xompass.png);
    background-repeat: no-repeat;
    background-size: contain;
}

.swagger-ui .topbar a img {
    display: none
}

.swagger-ui .topbar .download-url-wrapper {
    display: none;
}

.swagger-ui .info {
    margin: 30px 0;
}

.swagger-ui .info .title {
    font-size: 30px;
}

.swagger-ui .info .title small {
    border-radius: 3px;
}

.swagger-ui .info .title small pre {
    font-size: 10px;
}

.swagger-ui .scheme-container {
    padding: 15px 0;
}

.swagger-ui .scheme-container .servers, .swagger-ui .scheme-container .servers-title {
    display: none;
}

.swagger-ui .auth-wrapper .authorize {
    padding: 3px 20px;
}

.swagger-ui .opblock-tag {
    font-size: 18px;
    padding: 8px 20px 8px 8px;
}

.swagger-ui .opblock {
    border-radius: 0;
    margin-bottom: 10px;
}

.swagger-ui .opblock .opblock-summary {
    padding: 0;
}

.swagger-ui .opblock .opblock-summary-method {
    border-radius: 0;
    padding: 8px 15px;
    width: 100px;
}

.swagger-ui .opblock .opblock-summary-path {
    font-size: 14px;
}

.swagger-ui .opblock .opblock-summary-description {
    font-size: 12px;
    text-align: right;
}

.swagger-ui .parameter__name.required:after {
    display: none;
}

.swagger-ui .parameter__type,
.swagger-ui .parameter__deprecated,
.swagger-ui .parameter__in {
    font-size: 11px;
    padding-bottom: 0;
}

.swagger-ui table tbody tr td.parameters-col_name {
    min-width: 9em;
}

.swagger-ui table tbody tr td.parameters-col_description .renderedMarkdown p {
    margin: 0 0 5px;
    font-size: 14px;
}

.swagger-ui table tbody tr td.parameters-col_description input {
    margin-top: 0;
}

.swagger-ui table tbody tr td.parameters-col_description .renderedMarkdown.parameter__enum,
.swagger-ui table tbody tr td.parameters-col_description .renderedMarkdown.parameter__default {
    font-size: 13px;
    color: gray;
}

.swagger-ui section.models .model-container {
    border-radius: 0;
}

.swagger-ui section.models .model-container > .model-box {
    padding: 0;
}

.swagger-ui input[type=email],
.swagger-ui input[type=file],
.swagger-ui input[type=password],
.swagger-ui input[type=search],
.swagger-ui input[type=text] {
    padding: 5px 10px;
    font-size: 14px;
}

.swagger-ui textarea {
    padding: 5px 10px;
}

.swagger-ui .execute-wrapper {
    padding: 0 20px;
}

.swagger-ui .responses-wrapper {
    margin-top: 20px;
}`)
