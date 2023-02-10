
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def apikey = karate.properties['apikey']


Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"api-key": "#(apikey)",
		"model": "lambdal/text-to-pokemon",
		"version": "3554d9e699e09693d3fa334a79c58be9a405dd021d3e11281256d53185868912",
		"input": {
			"prompt": "a golden horse",
			"num_outputs": 1
		}
	}
	"""
	When method POST
	Then status 200





	# And path '/'
	# And header Direktiv-ActionID = 'development'
	# And header Direktiv-TempDir = '/tmp'
	# And request
	# """
	# {
	# 	"api-key": "#(apikey)",
	# 	"model": "lambdal/text-to-pokemon",
	# 	"version": "3554d9e699e09693d3fa334a79c58be9a405dd021d3e11281256d53185868912",
	# 	"input": {
	# 		"one": "two",
	# 		"three": "four",
	# 		"six": 6
	# 	},
	# 	"files": [
	# 		{
	# 			"name": "test",
	# 			"mime": "image/png",
	# 			"file": "/tests/in.png"
	# 		},
	# 		{
	# 			"name": "test1",
	# 			"mime": "image/png",
	# 			"file": "/tests/in.png"
	# 		}
	# 	]
	# }
	# """
	# When method POST
	# Then status 200