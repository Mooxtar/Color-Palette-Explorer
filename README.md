# Color-Palette-Explorer

Mukhtar Rabayev - Task 2 especially for NFactorial Incubator 2024

Color Palette Explorer is a web application that enables users to create, adjust, and save custom color palettes. It's built with Go, leveraging server-side rendering to deliver a dynamic, interactive experience.

## Features

Create Custom Palettes: Users can pick a base color and generate a color scheme.
Adjust Color Properties: Fine-tune hue, saturation, and lightness of individual colors.
Save(without database): Save color palettes.

Getting Started
To run Color Palette Explorer locally, you need Go installed on your machine. Follow these steps to get it up and running:

1. Clone the repository: git clone https://github.com/Mooxtar/color-palette-explorer.git
2. Navigate to the project directory: cd color-palette-explorer
3. Run the application: go run ./cmd/web

This will start the server on localhost:8080.


How It Works

User Interface:

The home.page.tmpl template provides a welcoming introduction to the application.
The most.page.tmpl template displays a gallery of the most used color palettes.
The generate.page.tmpl template allows users to generate new color palettes.

Color Scheme Generation:

Users can pick a base color which will be used to generate analogous color schemes.
The createColorScheme function in generate.page.tmpl dynamically creates color swatches based on the selected base color.

Adjusting Color Properties:

Users can adjust hue, saturation, and lightness of the selected color using the HTML range inputs.
The applyAdjustments function updates the color swatch to reflect the adjustments.

Saving Palettes:

Once satisfied with the created palette, users can save it.
The savePalette function clones the color scheme, assigns a unique identifier, and displays it under "Your Saved Palettes".


Notes: Tried to connect to database and save created color palettes to database, however the time was sufficient just to create db folder with database.go, which helps to connect to postgresql.

Thank you for your attention.


