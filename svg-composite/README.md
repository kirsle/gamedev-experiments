# SVG Compositor

Just a quick experiment about using SVG to design graphics for a game, export
the SVG to PNG files (at varying resolutions in production), and then the game
being able to compositely join the images together without any artifacting.

Use case: use Inkscape to draw the game graphics, including characters, to give
it a high quality cartoony look-and-feel (compile to PNG for performance, and
a final game would include the graphics compiled at different resolutions to
scale smoothly on different computers), and still allow for character
customization. For example, the game could dynamically rebuild character sprites
from pieces to have advanced customizability in-game.

This folder contains the original Inkscape SVG, three exported PNG components
(a face and two hairstyles), and a Python script that combines them and
produces two outputs (comp1.png and comp2.png).
