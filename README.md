# Project Voldemort

This repo exists to summon **You-Know-Who**, **He-Who-Shall-Not-Be-Named**, otherwise known as the **Dark Lord**.

It does not do anything useful.

## Pre-Requisites
- Obtain a facebox from [machinebox.io](https://machinebox.io/)
- Make sure the facebox is running

## Data Processing Pipeline

- Gather images of the person you would like to detect
- Run `voldemort data` to generate the training data in `training/data.json`
- Run `voldemort train` to teach the facebox how to find Tom Riddle.

## Face Detection

- Run `voldemort identify --image <URL>` to process the image
- Response JSON will contain matches

Example Response:
```
{
	"success": true,
	"facesCount": 2,
	"faces": [
		{
			"rect": {
				"top": 573,
				"left": 1294,
				"width": 186,
				"height": 185
			},
			"matched": false
		},
		{
			"rect": {
				"top": 201,
				"left": 386,
				"width": 186,
				"height": 186
			},
			"id": "59051e68e570abb39ece3c6395d72282",
			"name": "Tom Riddle",
			"matched": true
		}
	]
}
```

## Initialize Facebox State

- `.state` file can be used to initialize `MB_FACEBOX_STATE_URL`
