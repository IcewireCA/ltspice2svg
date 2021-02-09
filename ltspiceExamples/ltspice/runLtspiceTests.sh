#!/bin/bash
# shell script to backup this directory but ignore node_modules

ltspice2svg -export=all_symbols_default.svg all_symbols_default.asc

ltspice2svg -export=all_symbols_svg.svg all_symbols_svg.asc

ltspice2svg -export=ex_circ_block.svg ex_circ_block.asc

ltspice2svg -export=example.svg example.asc

ltspice2svg -export=exampleLatex.svg exampleLatex.asc

ltspice2svg -export=in_circuit.svg in_circuit.asc

ltspice2svg -export=in_draw.svg in_draw.asc

ltspice2svg -export=subscripts.svg subscripts.asc

ltspice2svg -export=wiresDotsJumpers.svg wiresDotsJumpers.asc

ltspice2svg -export=wiresNoDot.svg -Tdots=false wiresNoDot.asc