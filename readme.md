
# replicate 1.0

Generate images and videos with replicate.com

---
- #### Categories: misc
- #### Image: gcr.io/direktiv/functions/replicate 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/replicate/issues
- #### URL: https://github.com/direktiv-apps/replicate
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About replicate

This function can call the [replicate.com](https://replicate.com) API to generate images and videos with AI. 

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: replicate
  image: gcr.io/direktiv/functions/replicate:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: replicate
  type: action
  action:
    function: replicate
    secrets: ["replicate"]
    input:
      debug: true 
      model: stability-ai/stable-diffusion
      version: f178fa7a1ae43a9a9af01b833b9d2ecf97b1bcb0acfd2dc5dd04895e042863f1
      api-key: jq(.secrets.replicate)
      input:
        prompt: an green elephant jumping of a bridge with a parachute 
```
   #### Advanced
```yaml
- id: replicate
  type: action
  action:
    function: replicate
    input: 
      secrets: ["replicate"]
      model: lambdal/text-to-pokemon
      version: 3554d9e699e09693d3fa334a79c58be9a405dd021d3e11281256d53185868912
      input:
        one: two
        three: four
        six: 6
      files:
      - name: test
        mime: image/png
        file: "/tests/in.png"
  catch:
  - error: "*"
```

   ### Secrets


- **apikey**: Replicate.com API key






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Example Reponses
    
```json
{
  "completed_at": null,
  "created_at": "2023-02-10T09:28:22.977673Z",
  "error": null,
  "id": "qzk7dr2zjjasdsdsadada2ja4mmea",
  "input": {
    "num_outputs": 1,
    "prompt": "a golden horse"
  },
  "logs": "",
  "metrics": {},
  "output": null,
  "started_at": null,
  "status": "starting",
  "urls": {
    "cancel": "https://api.replicate.com/v1/predictions/qzk7dr2zjjasdsdsadada2ja4mmea/cancel",
    "get": "https://api.replicate.com/v1/predictions/qzk7dr2zjjasdsdsadada2ja4mmea"
  },
  "version": "3554d9e699e09693d3fa334a79c58be9a405dd021d3e11281256d53185868912",
  "webhook_completed": null
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| api-key | string| `string` | ✓ | | API key provided by replicate.com |  |
| debug | boolean| `bool` |  | | Print debug output |  |
| files | [][PostParamsBodyFilesItems](#post-params-body-files-items)| `[]*PostParamsBodyFilesItems` |  | |  |  |
| input | [interface{}](#interface)| `interface{}` |  | |  |  |
| model | string| `string` | ✓ | | Model for generation | `stability-ai/stable-diffusion` |
| version | string| `string` | ✓ | | Model version |  |


#### <span id="post-params-body-files-items"></span> postParamsBodyFilesItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| file | string| `string` | ✓ | | Path to the file provided by Direktiv |  |
| mime | string| `string` | ✓ | | Mimetype of the file | `image/png` |
| name | string| `string` | ✓ | | Name of the parameter in replicate.com API for the model |  |

 
