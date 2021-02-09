package main

func getDefaultSymDefn() string {
	// default symbol definitions for LTSpice
	var labelDefn = `
<defs>
	<g id="svg_supply">
      <polygon points="-6,0 0,-15 6,0" style="fill:black;fill-opacity:1;stroke:black;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
    <g id="svg_port">
      <circle cx="0" cy="0" r="6" style="stroke:black; fill:black"/>
      <circle cx="0" cy="0" r="3" style="stroke:white; fill:white"/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
    <g id="svg_dot">
      <circle cx="0" cy="0" r="5" style="stroke:#000000; fill:#000000"/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
    <g id="svg_jumpCirc">
      <circle cx="0" cy="0" r="6" style="stroke:#000000;stroke-width:2.5;fill:#ffffff"/>
      <rect x="-8" y="1.25" width="16" height="7" style="fill:#ffffff;stroke:none"/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
    <g id="svg_jumpLine">
      <path 
     d="M -12,0 12,0"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_gnd">
      <path 
     d="M 0,0 0,8
     M -14,8 14,8
     M -8,14 8,14
     M -2,20 2,20"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>

<defs>
     <g id="svg_block_adc">
    <path 
	d="
	M 32,56 80,56
	M 80,56 80,8
	M 80,8 32,8
	M 0,32 32,8
	M 0,32 32,56
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     <!-- default text positions and size 
			"svg_block_adc" WINDOW 3 48 32 Center 2
			"svg_block_adc" SYMATTR Value ADC
		-->
     </g>
</defs>
<defs>
     <g id="svg_block_amp">
      <polygon points="16,-16 112,32 16,80" style="fill:none;fill-opacity:1;stroke:black;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     <!-- default text positions and size 
			"svg_block_amp" WINDOW 3 48 32 Center 2
			"svg_block_amp" SYMATTR Value text
		-->
     </g>
</defs>
<defs>
     <g id="svg_block_arrow">
      <polygon points="-16,-5 -16,5 0,0" style="fill:black;fill-opacity:1;stroke:black;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_block_arrow2">
     <path
    d="
	M 0,0 -17,-13
	M -17,13 0,0
	M 0,0 -16,0
    "
    style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_block_box">
    <path 
	d="
	M 16,-8 16,72
	M 16,72 112,72
	M 112,72 112,-8
	M 112,-8 16,-8
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     <!-- default text positions and size 
			"svg_block_box" WINDOW 3 64 32 Center 2
			"svg_block_box" SYMATTR Value text
		-->
     </g>
</defs>
<defs>
     <g id="svg_block_dac">
    <path 
	d="
	M 48,56 0,56
	M 0,56 0,8
	M 0,8 48,8
	M 80,32 48,8
	M 80,32 48,56
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     <!-- default text positions and size 
			"svg_block_dac" WINDOW 3 32 32 Center 2
			"svg_block_dac" SYMATTR Value DAC
		-->
     </g>
</defs>
<defs>
     <g id="svg_block_delay">
    <path 
	d="
	M 0,8 0,56
	M 0,56 64,56
	M 64,56 64,8
	M 64,8 0,8
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<text x="28" y="36" style="white-space:pre;dominant-baseline:middle;text-anchor:middle;font-size:26px;">z</text>
	<text x="44" y="24" style="white-space:pre;dominant-baseline:middle;text-anchor:middle;font-size:20px;">-1</text>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_block_quant">
    <path 
	d="
	M 16,8 16,56
	M 16,56 80,56
	M 80,56 80,8
	M 80,8 16,8
	M 24,48 40,48
	M 40,48 40,32
	M 40,32 56,32
	M 56,32 56,16
	M 56,16 72,16
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_block_sum">
     <circle cx="0" cy="16" r="16" style="stroke:black; stroke-width:2.5;fill:white"/>
      <path 
	d="M 0,8 0,24
	M 8,16 -8,16"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
	  <g id="svg_cap">
		  <path 
		 d="M 16,37 16,64
		 M 16,27 16,0" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <path 
		 d="M -4,27 36,27
		 M -4,37 36,37" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		   <!-- default text positions and size 
		 	"svg_cap" WINDOW 0 44 18 Left 2
			"svg_cap" WINDOW 3 44 46 Left 2
			"svg_cap" SYMATTR Value C
			"svg_cap" SYMATTR Prefix C
		   -->
	  </g>
	</defs>
<defs>
	<g id="svg_current">
	<circle cx="0" cy="32" r="24" style="stroke:black; stroke-width:4;fill:white"/>
	<path 
	d="
	M 0,48 4,36
	M 0,48 -4,36
	M -4,36 4,36
	M 0,16 0,36
	M 0,64 0,56
	M 0,0 0,8
	"
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size 
		"svg_current" WINDOW 0 -36 20 Right 2
		"svg_current" WINDOW 3 -36 44 Right 2
		"svg_current" SYMATTR Value I
		"svg_current" SYMATTR Prefix I
	-->
	</g>
</defs>
<defs>
	<g id="svg_diode">
	<path 
	d="
	M 16,0 16,20
	M 16,44 16,64
	"
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<path 
	d="
	M 0,44 32,44
	M 0,20 32,20
	M 32,20 16,44
	M 0,20 16,44
	"
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size 
		"svg_diode" WINDOW 0 40 20 Left 2
		"svg_diode" WINDOW 3 40 44 Left 2
		"svg_diode" SYMATTR Value D
		"svg_diode" SYMATTR Prefix D
	-->
	</g>
</defs>
<defs>
	  <g id="svg_ind">
	  	<path 
		 d="
		 M 16,0 16,8 
		 M 16,56 16,64
		 "
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		 />
		 <path 
		 d="
		 M 16,8 C 44,8 44,24 16,24
		 M 16,24 C 44,24 44,40 16,40
		 M 16,40 C 44,40 44,56 16,56 
		 "
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		 />
		 <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		 <!-- default text positions and size 
		 	"svg_ind" WINDOW 0 40 18 Left 2
			"svg_ind" WINDOW 3 40 46 Left 2
			"svg_ind" SYMATTR Value L
			"svg_ind" SYMATTR Prefix L
		-->
	  </g>
</defs>
<defs>
     <g id="svg_meas_I">
     <path 
     d="M 32,0 32,64
     M 32,0 40,16
     M 32,0 24,16"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_meas_imped">
     <path 
     d="
     M 32,0 32,36
	M 32,0 40,16
	M 32,0 24,16
	M 32,36 4,36
"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_meas_minus">
     <path 
     d="
     M 16,0 32,0
     "
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_meas_plus">
     <path 
     d="
     M 16,0 32,0
     M 24,-8 24,8
     "
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_meas_V">
     <path 
     d="
     M 16,0 32,0
     M 24,-8 24,8
     M 32,64 16,64
     "
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_meas_V2">
     <path 
     d="
     M 24,0 40,0
     M 32,-8 32,8
     M -40,0 -24,0
     "
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
     <g id="svg_meas_V3">
     <path 
     d="
     M 16,0 32,0
     M 24,-8 24,8
     M 64,32 80,32
     "
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
<g id="svg_meas_Vminus">
<polyline points="16 0 32 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"svg_meas_Vminus" SYMATTR Description Voltage +-
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="svg_meas_Vplus">
<polyline points="16 0 32 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 -8 24 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"svg_meas_Vplus" SYMATTR Description Voltage +-
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
	  <g id="svg_nmos">
		  <path 
		 d="M 0,48 16,48
		 M 64,28 28,28
		 M 64,68 28,68
		 M 64,68 64,80
		 M 64,28 64,16" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <path 
		 d="M 16,70 16,24
		 M 28,81 28,15" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:6;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <polygon points="47,59 47,77 60,68" style="fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
		   <!-- Below is an invisible box so that the symbol stays on grid 
		   when in inkscape -->
		   <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		   <!-- default text positions and size
		 	"svg_nmos" WINDOW 0 70 48 Left 2
		 	"svg_nmos" SYMATTR Prefix MN
		   -->
	  </g>
	</defs>
<defs>
	  <g id="svg_nmos4">
		  <path 
		 d="M 0,48 16,48
		 M 64,28 28,28
		 M 64,68 28,68
		 M 64,68 64,80
		 M 64,28 64,16
		 M 64,48 28,48
		 " 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <path 
		 d="M 16,70 16,24
		 M 28,81 28,15" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:6;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <polygon points="47,59 47,77 60,68" style="fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
		   <!-- Below is an invisible box so that the symbol stays on grid 
		   when in inkscape -->
		   <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		   <!-- default text positions and size
		 	"svg_nmos4" WINDOW 0 70 32 Left 2
		 	"svg_nmos4" SYMATTR Prefix MN
		   -->
	  </g>
	</defs>
<defs>
	<g id="svg_npn">
	<path 
	d="
	M 48,60 44,72
	M 64,72 48,60
	M 64,72 44,72
	M 46,66 24,56
	M 24,72 24,24
	M 24,40 64,24
	M 24,48 0,48
	M 64,72 64,80
	M 64,16 64,24
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<!-- Below is an invisible box so that the symbol stays on grid 
	when in inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size
	"svg_npn" WINDOW 0 70 48 Left 2
	"svg_npn" SYMATTR Prefix MN
	-->
	</g>
</defs>
<defs>
<g id="svg_npn1">
<polyline points="48 64 40 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 80 48 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 80 40 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="44 70 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 64 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"svg_npn1" WINDOW 0 70 48 Left 2
"svg_npn1" SYMATTR Value NPN
"svg_npn1" SYMATTR Prefix QN
"svg_npn1" SYMATTR Description NPN
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
     <g id="svg_opamp">
     <path 
     d="
	M -28,48 -20,48
	M -28,80 -20,80
	M -24,84 -24,76
"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     /> 
     <path 
     d="
    M -32,32 32,64
	M -32,96 32,64
	M -32,32 -32,96
"
     style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
     />      
     <rect x="-48" y="-48" width="208" height="208" style="fill:none;stroke:none" />
     </g>
</defs>
<defs>
	  <g id="svg_pmos">
		  <path 
		 d="M 0,48 16,48
		 M 64,28 28,28
		 M 64,68 28,68
		 M 64,68 64,80
		 M 64,28 64,16" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <path 
		 d="M 16,70 16,24
		 M 28,81 28,15" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:6;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <polygon points="45,19 45,37 32,28" style="fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
		   <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		   <!-- default text positions and size
		 	"svg_pmos" WINDOW 0 70 48 Left 2
		 	"svg_pmos" SYMATTR Prefix MP
		   -->
	  </g>
	</defs>
<defs>
	  <g id="svg_pmos4">
		  <path 
		 d="M 0,48 16,48
		 M 64,28 28,28
		 M 64,68 28,68
		 M 64,68 64,80
		 M 64,28 64,16
		 M 64,48 28,48
		 " 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <path 
		 d="M 16,70 16,24
		 M 28,81 28,15" 
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:6;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		   />
		   <polygon points="45,19 45,37 32,28" style="fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
		   <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		   <!-- default text positions and size
		 	"svg_pmos4" WINDOW 0 70 70 Left 2
		 	"svg_pmos4" SYMATTR Prefix MP
		   -->
	  </g>
	</defs>
<defs>
	<g id="svg_pnp">
	<path 
	d="
	M 40,28 44,40
	M 24,40 40,28
	M 24,40 44,40
	M 64,72 24,56
	M 24,72 24,24
	M 42,34 64,24
	M 24,48 0,48
	M 64,72 64,80
	M 64,16 64,24
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<!-- Below is an invisible box so that the symbol stays on grid 
	when in inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size
	"svg_pnp" WINDOW 0 70 48 Left 2
	"svg_pnp" SYMATTR Prefix MN
	-->
	</g>
</defs>
<defs>
<g id="svg_pnp1">
<polyline points="16 58 64 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="35 31 64 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 38 37" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="38 37 32 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 32 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"svg_pnp1" WINDOW 0 70 48 Left 2
"svg_pnp1" SYMATTR Value PNP
"svg_pnp1" SYMATTR Prefix QP
"svg_pnp1" SYMATTR Description PNP
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
	  <g id="svg_res">
		  <path 
		 d="M 16,0 32,5
		 M 32,5 0,16
		 M 0,16 32,27
		 M 32,27 0,37
		 M 0,37 32,48
		 M 32,48 0,59
		 M 0,59 16,64"
		 style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
		 />
		 <rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		 <!-- default text positions and size 
		 	"svg_res" WINDOW 0 40 18 Left 2 
			"svg_res" WINDOW 3 40 46 Left 2
			"svg_res" SYMATTR Value R
			"svg_res" SYMATTR Prefix R
		-->
	  </g>
</defs>
<defs>
<g id="svg_ss_nmos">
<polyline points="16 70 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 81 24 15" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="55 68 24 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="55 63 55 73" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 68 55 63" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="55 73 64 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 28 64 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 28 64 26" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 68 64 70" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 10 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 25.00 A 5.00 2.00 0 0 0 64.00 21.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 16.00 A 5.00 2.00 0 0 0 64.00 20.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 74.00 A 5.00 2.00 0 0 0 64.00 70.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 75.00 A 5.00 2.00 0 0 0 64.00 79.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 5.00 48.00 A 2.00 5.00 0 0 0 9.00 48.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 4.00 48.00 A 2.00 5.00 0 0 0 0.00 48.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"svg_ss_nmos" WINDOW 0 70 48 Left 2
"svg_ss_nmos" SYMATTR Description Small signal NMOS
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="svg_ss_pmos">
<polyline points="16 74 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 80 24 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 68 24 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 28 33 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="33 23 33 33" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 28 33 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="33 33 24 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 10 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 68 64 70" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 28 64 26" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 25.00 A 5.00 2.00 0 0 0 64.00 21.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 16.00 A 5.00 2.00 0 0 0 64.00 20.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 74.00 A 5.00 2.00 0 0 0 64.00 70.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 64.00 75.00 A 5.00 2.00 0 0 0 64.00 79.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 5.00 48.00 A 2.00 5.00 0 0 0 9.00 48.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 4.00 48.00 A 2.00 5.00 0 0 0 0.00 48.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"svg_ss_pmos" WINDOW 0 65 48 Left 2
"svg_ss_pmos" SYMATTR Description Small Signal PMOS
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
	<g id="svg_vccs">
	<path 
	d="
	M 0,48 4,36
	M 0,48 -4,36
	M -4,36 4,36
	M 0,16 0,36
	M 0,64 0,56
	M 0,0 0,8
	M 0,56 24,32
	M 24,32 0,8
	M 0,8 -24,32
	M -24,32 0,56
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<!-- Below is an invisible box so that the symbol stays on grid 
	when in inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size
	"svg_vccs" WINDOW 3 32 32 Left 2
	"svg_vccs" SYMATTR Value GmVx
	-->
	</g>
</defs>
<defs>
	<g id="svg_vcvs">
	<path 
	d="
	M -8,24 8,24
	M -8,40 8,40
	M 0,16 0,32
	M 0,64 0,56
	M 0,0 0,8
	M 0,8 24,32
	M 24,32 0,56
	M 0,56 -24,32
	M -24,32 0,8
	" 
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<!-- Below is an invisible box so that the symbol stays on grid 
	when in inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size
	"svg_vcvs" WINDOW 3 32 32 Left 2
	"svg_vcvs" SYMATTR Value KVx
	-->
	</g>
</defs>
<defs>
	<g id="svg_voltage">
	<circle cx="0" cy="32" r="24" style="stroke:black; stroke-width:4;fill:white"/>
	<path 
	d="
	M -8,24 8,24
	M -8,44 8,44
	M 0,16 0,32
	M 0,64 0,56
	M 0,0 0,8
	"
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size 
		"svg_voltage" WINDOW 0 -36 20 Right 2
		"svg_voltage" WINDOW 3 -36 44 Right 2
		"svg_voltage" SYMATTR Value I
		"svg_voltage" SYMATTR Prefix I
	-->
	</g>
</defs>
<defs>
	<g id="svg_voltage2">
	<circle cx="0" cy="32" r="24" style="stroke:black; stroke-width:4;fill:white"/>
	<path 
	d="
	M 0,64 0,56
	M 0,0 0,8
	"
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	<g transform = "translate(-16, 32)">
	<path 
	d="M 0,0 C 1.23347,-2.9265965 2.4669399,-5.8412165 3.7004097,-8.0837765 4.9338798,-10.326335 6.1673494,-11.875859 7.4008195,-12.385062 8.6342896,-12.894267 9.8677596,-12.353654 11.10123,-10.891241 12.3347,-9.4288287 13.568169,-7.0510391 14.80164,-4.3012874 16.03511,-1.5515357 17.26858,1.550837 18.502049,4.3012875 19.73552,7.0517382 20.968989,9.4270602 22.20246,10.891243 23.435929,12.355426 24.669398,12.892256 25.902869,12.385064 27.136339,11.877872 28.369809,10.325023 29.603279,8.0837762 30.836749,5.8425298 32.070219,2.9265966 33.303687,0"
	style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
	/>
	</g>
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	<!-- default text positions and size 
		"svg_voltage2" WINDOW 0 -36 20 Right 2
		"svg_voltage2" WINDOW 3 -36 44 Right 2
		"svg_voltage2" SYMATTR Value I
		"svg_voltage2" SYMATTR Prefix I
	-->
	</g>
</defs>

<defs>
<g id="bi">
<polyline points="0 56 4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 -4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 44 4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 24 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"bi" WINDOW 0 24 0 Left 2
"bi" WINDOW 3 24 80 Left 2
"bi" SYMATTR Value I=F(...)
"bi" SYMATTR Prefix B
"bi" SYMATTR Description Arbitrary behavioral current source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="bi2">
<polyline points="0 24 4 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 24 -4 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 36 4 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 36 0 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"bi2" WINDOW 0 24 0 Left 2
"bi2" WINDOW 3 24 81 Left 2
"bi2" SYMATTR Value I=F(...)
"bi2" SYMATTR Prefix B
"bi2" SYMATTR Description Arbitrary behavioral current source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="bv">
<polyline points="-8 36 8 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 76 8 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"bv" WINDOW 0 24 16 Left 2
"bv" WINDOW 3 24 96 Left 2
"bv" SYMATTR Value V=F(...)
"bv" SYMATTR Prefix B
"bv" SYMATTR Description Arbitrary behavioral voltage source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="cap">
<polyline points="16 36 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 28 16 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 32 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 36 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"cap" WINDOW 0 24 8 Left 2
"cap" WINDOW 3 24 56 Left 2
"cap" SYMATTR Value C
"cap" SYMATTR Prefix C
"cap" SYMATTR Description Capacitor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="csw">
<polyline points="0 80 0 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 20 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="20" cy="44" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"csw" WINDOW 0 24 0 Left 2
"csw" WINDOW 3 24 88 Left 2
"csw" SYMATTR Value CSW
"csw" SYMATTR Prefix W
"csw" SYMATTR Description Current controlled switch
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="current">
<polyline points="0 56 4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 -4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 44 4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 24 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"current" WINDOW 0 24 0 Left 2
"current" WINDOW 3 24 80 Left 2
"current" SYMATTR Value I
"current" SYMATTR Prefix I
"current" SYMATTR Description Current source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="diode">
<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"diode" WINDOW 0 24 0 Left 2
"diode" WINDOW 3 24 64 Left 2
"diode" SYMATTR Value D
"diode" SYMATTR Prefix D
"diode" SYMATTR Description Diode
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="e">
<polyline points="-48 32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 32 -24 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 80 -32 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 80 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 72 -40 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 40 -40 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-44 36 -44 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 72 4 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 40 4 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 36 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"e" WINDOW 0 24 16 Left 2
"e" WINDOW 3 24 96 Left 2
"e" SYMATTR Value E
"e" SYMATTR Prefix E
"e" SYMATTR Description Voltage dependent voltage source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="e2">
<polyline points="-48 32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 32 -24 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 80 -32 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 80 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 72 -40 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-44 76 -44 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 40 -40 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 72 4 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 40 4 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 36 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"e2" WINDOW 0 24 16 Left 2
"e2" WINDOW 3 24 96 Left 2
"e2" SYMATTR Value E
"e2" SYMATTR Prefix E
"e2" SYMATTR Description Voltage dependent voltage source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="f">
<polyline points="0 56 4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 -4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 44 4 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 24 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"f" WINDOW 0 24 0 Left 2
"f" WINDOW 3 24 80 Left 2
"f" SYMATTR Value F
"f" SYMATTR Prefix F
"f" SYMATTR Description Linear current dependent current source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="FerriteBead">
<polyline points="-16 8 -16 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 -8 0 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 12 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 8 16 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="-8" rx="16" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="-8" rx="7" ry="1" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M -16.00 8.00 A 16.00 4.00 0 0 0 16.00 8.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"FerriteBead" WINDOW 0 25 0 Left 2
"FerriteBead" SYMATTR Prefix L_Ferrite_Bead
"FerriteBead" SYMATTR Description A Ferrite Bead
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="FerriteBead2">
<polyline points="0 -9 0 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 9 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 8 9 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 7 9 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 6 9 6" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 6 9 6" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 5 9 5" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 4 9 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 3 9 3" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 3 9 3" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 2 9 2" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 1 9 1" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 0 9 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 0 9 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -1 9 -1" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -4 9 -4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -5 9 -5" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -6 9 -6" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -6 9 -6" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -7 9 -7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -8 9 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -3 9 -3" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-9 -2 9 -2" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="7 9 7 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 9 8 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="5 9 5 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="6 9 6 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="3 9 3 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 9 4 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="1 9 1 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="2 9 2 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-1 9 -1 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 9 0 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-3 9 -3 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-2 9 -2 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-5 9 -5 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 9 -4 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-7 9 -7 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-6 9 -6 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 9 -8 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<rect x="-9" y="-9" width="18" height="18" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"FerriteBead2" WINDOW 0 21 0 Left 2
"FerriteBead2" SYMATTR Prefix L_Ferrite_Bead
"FerriteBead2" SYMATTR Description A Ferrite Bead(Alternate symbol)
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="g">
<polyline points="-48 32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 32 -24 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 80 -32 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 80 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 72 -40 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 40 -40 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-44 36 -44 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 52 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 52 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 52 4 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 52 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"g" WINDOW 0 24 16 Left 2
"g" WINDOW 3 24 96 Left 2
"g" SYMATTR Value G
"g" SYMATTR Prefix G
"g" SYMATTR Description Voltage dependent current source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="g2">
<polyline points="-48 32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 32 -24 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 80 -32 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 80 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 72 -40 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 40 -40 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-44 76 -44 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 52 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 52 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 52 4 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 52 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"g2" WINDOW 0 24 16 Left 2
"g2" WINDOW 3 24 96 Left 2
"g2" SYMATTR Value G
"g2" SYMATTR Prefix G
"g2" SYMATTR Description Voltage dependent current source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="h">
<polyline points="-8 36 8 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 76 8 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"h" WINDOW 0 24 16 Left 2
"h" WINDOW 3 24 96 Left 2
"h" SYMATTR Value H
"h" SYMATTR Prefix H
"h" SYMATTR Description Linear current dependent voltage source
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="ind">
<path d=" M 4.69 67.31 A 16.00 16.00 0 1 0 4.69 44.69" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 16.00 96.00 A 16.00 16.00 0 1 0 4.69 68.69" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 4.69 43.31 A 16.00 16.00 0 1 0 16.00 16.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"ind" WINDOW 0 36 40 Left 2
"ind" WINDOW 3 36 80 Left 2
"ind" SYMATTR Value L
"ind" SYMATTR Prefix L
"ind" SYMATTR Description Inductor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="ind2">
<ellipse cx="8" cy="84" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 4.69 67.31 A 16.00 16.00 0 1 0 4.69 44.69" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 16.00 96.00 A 16.00 16.00 0 1 0 4.69 68.69" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 4.69 43.31 A 16.00 16.00 0 1 0 16.00 16.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"ind2" WINDOW 0 36 40 Left 2
"ind2" WINDOW 3 36 80 Left 2
"ind2" SYMATTR Value L
"ind2" SYMATTR Prefix L
"ind2" SYMATTR Description Inductor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="ISO16750-2">
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 55 -8 18" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-18 55 -8 55" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M -8.00 18.00 A 29.00 36.00 0 0 0 20.22 53.99" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"ISO16750-2" WINDOW 0 48 16 Left 2
"ISO16750-2" WINDOW 38 48 64 Left 2
"ISO16750-2" WINDOW 1 48 40 Left 2
"ISO16750-2" SYMATTR SpiceModel 4-6-3_12V_StartingProfile
"ISO16750-2" SYMATTR Description ISO-16750-2 Transients
"ISO16750-2" SYMATTR Prefix X
"ISO16750-2" SYMATTR ModelFile ISO16750-2.lib
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="ISO7637-2">
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 55 -8 18" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-18 55 -8 55" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M -8.00 18.00 A 29.00 36.00 0 0 0 20.22 53.99" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"ISO7637-2" WINDOW 0 48 16 Left 2
"ISO7637-2" WINDOW 38 48 64 Left 2
"ISO7637-2" WINDOW 1 48 40 Left 2
"ISO7637-2" SYMATTR SpiceModel Pulse1_12V
"ISO7637-2" SYMATTR Description ISO-7637-2 Transeints
"ISO7637-2" SYMATTR Prefix X
"ISO7637-2" SYMATTR ModelFile ISO7637-2.lib
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="LED">
<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 32 68 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 32 64 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 48 68 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 48 64 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 56.00 28.00 A 8.00 8.00 0 0 0 40.84 24.42" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 56.00 28.00 A 8.00 8.00 0 0 0 71.16 31.58" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 56.00 44.00 A 8.00 8.00 0 0 0 40.84 40.42" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 56.00 44.00 A 8.00 8.00 0 0 0 71.16 47.58" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"LED" WINDOW 0 24 0 Left 2
"LED" WINDOW 3 24 64 Left 2
"LED" SYMATTR Value D
"LED" SYMATTR Prefix D
"LED" SYMATTR Description Light Emitting Diode
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="load">
<polyline points="0 48 40 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 8 28 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 8 36 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 8 8 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 56 24 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 56 24 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 8 8 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 56 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"load" WINDOW 0 48 8 Left 2
"load" WINDOW 3 48 56 Left 2
"load" SYMATTR Value I
"load" SYMATTR Prefix I
"load" SYMATTR Description Current Source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="load2">
<polyline points="36 4 32 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="36 4 24 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 72 36 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 4 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 -4 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-4 48 4 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 24 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 0 0 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="40" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"load2" WINDOW 0 17 -16 Left 2
"load2" WINDOW 3 25 84 Left 2
"load2" SYMATTR Value I
"load2" SYMATTR Prefix I
"load2" SYMATTR Description Current Source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="lpnp">
<polyline points="16 64 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 64 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 64 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="52 48 64 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 36 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="36 40 36 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="52 40 52 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="52 40 36 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="52 56 36 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"lpnp" WINDOW 0 72 32 Left 2
"lpnp" WINDOW 3 72 68 Left 2
"lpnp" SYMATTR Value LPNP
"lpnp" SYMATTR Prefix QP
"lpnp" SYMATTR Description Bipolar Lateral PNP transistor.  You would normally use this with your own .model statement.
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="ltline">
<polyline points="-48 16 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-36 12 36 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-36 -12 36 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -16 48 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 16 -36 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -16 -36 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 36 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -16 36 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-40 0 -32 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 0 -24 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-24 8 -8 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 -8 8 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 8 24 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 -8 32 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 0 40 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"ltline" WINDOW 0 0 -16 Bottom 2
"ltline" WINDOW 3 0 16 Top 2
"ltline" SYMATTR Value LTRA
"ltline" SYMATTR Prefix O
"ltline" SYMATTR Description Lossy Transmission Line.  Note: You must supply the .model for this device
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="mesfet">
<polyline points="16 12 16 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 80 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 48 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 16 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 16 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 48 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"mesfet" WINDOW 0 56 32 Left 2
"mesfet" WINDOW 3 56 72 Left 2
"mesfet" SYMATTR Value NMF
"mesfet" SYMATTR Prefix Z
"mesfet" SYMATTR Description GaAs MESFET transistor  Note: You must supply the .model for this device
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="njf">
<polyline points="16 16 16 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 72 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 48 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 24 48 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 24 48 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 64 4 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 68 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 60 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 60 4 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"njf" WINDOW 0 56 32 Left 2
"njf" WINDOW 3 56 72 Left 2
"njf" SYMATTR Value NJF
"njf" SYMATTR Prefix JN
"njf" SYMATTR Description N-Channel JFET
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="nmos">
<polyline points="48 48 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 48 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 48 48 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 40 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 40 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 44 40 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 8 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 16 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 0 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"nmos" WINDOW 0 56 32 Left 2
"nmos" WINDOW 3 56 72 Left 2
"nmos" SYMATTR Value NMOS
"nmos" SYMATTR Prefix MN
"nmos" SYMATTR Description N-Channel MOSFET transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="nmos4">
<polyline points="48 80 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 48 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 48 48 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 40 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 40 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 44 40 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 8 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 16 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 0 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"nmos4" WINDOW 0 56 32 Left 2
"nmos4" WINDOW 3 56 72 Left 2
"nmos4" SYMATTR Value NMOS
"nmos4" SYMATTR Prefix MN
"nmos4" SYMATTR Description N-Channel MOSFET transistor with explicit substrate connection(used for monolithic MOSFETS)
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="npn">
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"npn" WINDOW 0 56 32 Left 2
"npn" WINDOW 3 56 68 Left 2
"npn" SYMATTR Value NPN
"npn" SYMATTR Prefix QN
"npn" SYMATTR Description Bipolar NPN transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="npn2">
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<rect x="12" y="24" width="4" height="48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"npn2" WINDOW 0 56 32 Left 2
"npn2" WINDOW 3 56 68 Left 2
"npn2" SYMATTR Value NPN
"npn2" SYMATTR Prefix QN
"npn2" SYMATTR Description Bipolar NPN transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="npn3">
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 20 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="20 72 20 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 72 12 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="20 72 12 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="20 24 12 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="20 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"npn3" WINDOW 0 56 32 Left 2
"npn3" WINDOW 3 56 68 Left 2
"npn3" SYMATTR Value NPN
"npn3" SYMATTR Prefix QN
"npn3" SYMATTR Description Bipolar NPN transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="npn4">
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 96 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 16 48 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 48 56 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 32 56 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 44 64 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 44 48 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 32 48 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"npn4" WINDOW 0 72 32 Left 2
"npn4" WINDOW 3 72 68 Left 2
"npn4" SYMATTR Value NPN
"npn4" SYMATTR Prefix QN
"npn4" SYMATTR Description Bipolar NPN transistor with substrate node.  You would normally use this with your own .model statement.
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="pjf">
<polyline points="16 16 16 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 72 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 48 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 24 48 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 24 48 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 64 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 68 0 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 60 0 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 60 12 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"pjf" WINDOW 0 56 32 Left 2
"pjf" WINDOW 3 56 72 Left 2
"pjf" SYMATTR Value PJF
"pjf" SYMATTR Prefix JP
"pjf" SYMATTR Description P-Channel JFET transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="pmos">
<polyline points="48 48 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 48 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 24 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 48 24 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 48 24 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 44 24 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 8 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 16 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 0 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"pmos" WINDOW 0 56 32 Left 2
"pmos" WINDOW 3 56 72 Left 2
"pmos" SYMATTR Value PMOS
"pmos" SYMATTR Prefix MP
"pmos" SYMATTR Description P-Channel MOSFET transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="pmos4">
<polyline points="48 80 48 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 48 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 24 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 48 24 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 48 24 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="24 44 24 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 8 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 40 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 72 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 16 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 0 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"pmos4" WINDOW 0 56 32 Left 2
"pmos4" WINDOW 3 56 72 Left 2
"pmos4" SYMATTR Value PMOS
"pmos4" SYMATTR Prefix MP
"pmos4" SYMATTR Description P-Channel MOSFET transistor with explicit substrate connection(used for monolithic MOSFETS)
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="pnp">
<polyline points="16 64 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 64 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 64 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"pnp" WINDOW 0 84 32 Left 2
"pnp" WINDOW 3 84 68 Left 2
"pnp" SYMATTR Value PNP
"pnp" SYMATTR Prefix QP
"pnp" SYMATTR Description Bipolar PNP transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="pnp2">
<polyline points="16 64 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 64 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 64 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="12 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<rect x="12" y="24" width="4" height="48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"pnp2" WINDOW 0 84 32 Left 2
"pnp2" WINDOW 3 84 68 Left 2
"pnp2" SYMATTR Value PNP
"pnp2" SYMATTR Prefix QP
"pnp2" SYMATTR Description Bipolar PNP transistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="pnp4">
<polyline points="16 64 44 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="44 76 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 64 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 80 64 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 80 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 32 64 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 48 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 16 48 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="64 48 56 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 32 56 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 44 64 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 32 56 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="56 20 56 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"pnp4" WINDOW 0 72 32 Left 2
"pnp4" WINDOW 3 72 68 Left 2
"pnp4" SYMATTR Value PNP
"pnp4" SYMATTR Prefix QP
"pnp4" SYMATTR Description Bipolar PNP transistor with substrate node.  You would normally use this with your own .model statement.
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="polcap">
<polyline points="16 36 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="8 12 8 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="4 16 12 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 32 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<path d=" M 31.88 40.22 A 32.00 32.00 0 0 0 0.12 40.22" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"polcap" WINDOW 0 24 8 Left 2
"polcap" WINDOW 3 24 57 Left 2
"polcap" SYMATTR Value C
"polcap" SYMATTR Prefix C
"polcap" SYMATTR Description Polarized Capacitor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="res">
<polyline points="16 88 16 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 80 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 64 0 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 48 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 32 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 16 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 24 32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"res" WINDOW 0 36 40 Left 2
"res" WINDOW 3 36 76 Left 2
"res" SYMATTR Value R
"res" SYMATTR Prefix R
"res" SYMATTR Description A resistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="res2">
<polyline points="16 0 32 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 12 32 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 12 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 44 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 44 32 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 60 32 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 60 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"res2" WINDOW 0 36 16 Left 2
"res2" WINDOW 3 36 56 Left 2
"res2" SYMATTR Value R
"res2" SYMATTR Prefix R
"res2" SYMATTR Description A resistor
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="schottky">
<polyline points="0 36 4 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 44 0 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 44 32 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 52 28 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"schottky" WINDOW 0 24 0 Left 2
"schottky" WINDOW 3 24 64 Left 2
"schottky" SYMATTR Value D
"schottky" SYMATTR Prefix D
"schottky" SYMATTR Description Schottky diode
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="SOAtherm-HeatSink">
<polyline points="48 64 -48 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -7 32 -7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -16 48 -7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -16 48 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -40 32 -40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -48 48 -40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -64 -48 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -48 48 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -56 32 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -56 32 -56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -64 48 -56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-80 64 -48 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:6 6;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-80 -64 -80 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:6 6;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -64 -80 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:6 6;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 64 -48 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -24 32 -24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -32 48 -24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -32 48 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -24 32 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -40 32 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 55 48 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 55 48 55" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 31 32 31" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 23 48 31" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 23 48 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 15 32 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 15 32 15" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 7 48 15" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 47 32 47" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 39 48 47" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 39 48 39" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 47 32 55" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 31 32 39" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 7 32 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 -7 32 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -96 -80 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:6 6;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-16 -96 -48 -96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:6 6;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -64 -16 -96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -96 -16 -96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -64 80 -96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -88 48 -56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -80 48 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -72 48 -40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -64 48 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -56 48 -24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -48 48 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -39 48 -7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -25 48 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -17 48 15" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -9 48 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -1 48 31" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 7 48 39" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 15 48 47" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 23 48 55" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 32 48 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -88 80 -96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -72 80 -80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -56 80 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -39 80 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -17 80 -25" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 -1 80 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 15 80 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 32 80 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 -56 32 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 -40 32 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 -24 32 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="46 -7 32 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 15 32 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 31 32 39" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="40 47 32 55" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 -80 80 -80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 -64 80 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 -48 80 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 23 80 23" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 7 80 7" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="72 -9 80 -9" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="66 -25 80 -25" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"SOAtherm-HeatSink" WINDOW 0 0 -64 Bottom 2
"SOAtherm-HeatSink" WINDOW 3 97 29 Left 2
"SOAtherm-HeatSink" WINDOW 123 97 52 Left 2
"SOAtherm-HeatSink" WINDOW 38 97 -28 Left 2
"SOAtherm-HeatSink" SYMATTR Value Rtheta=10 Tambient=85
"SOAtherm-HeatSink" SYMATTR Value2 Area_Contact_mm2=100 Volume_mm3=1000
"SOAtherm-HeatSink" SYMATTR SpiceModel copper
"SOAtherm-HeatSink" SYMATTR Prefix X
"SOAtherm-HeatSink" SYMATTR ModelFile SOAtherm-HeatSink.lib
"SOAtherm-HeatSink" SYMATTR Description Heat Sink Thermal Model
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="SOAtherm-PCB">
<polyline points="-79 -40 -112 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="112 -40 -79 -40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 40 112 -40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-112 40 80 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-112 48 -112 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 48 -112 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="80 48 80 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="112 -32 80 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="112 -40 112 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:4 4;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"SOAtherm-PCB" WINDOW 0 16 -40 Bottom 2
"SOAtherm-PCB" WINDOW 3 -112 64 Left 2
"SOAtherm-PCB" WINDOW 123 -112 96 Left 2
"SOAtherm-PCB" SYMATTR Value Area_Contact_mm2=100 Area_PCB_mm2=2500 Copper_Thickness_oz=1
"SOAtherm-PCB" SYMATTR Value2 Tambient=85 LFM=100 PCB_FR4_Thickness_mm=1.5
"SOAtherm-PCB" SYMATTR Prefix X
"SOAtherm-PCB" SYMATTR SpiceModel TopsideCopper
"SOAtherm-PCB" SYMATTR ModelFile SOAtherm-PCB.lib
"SOAtherm-PCB" SYMATTR Description PCB Thermal Model
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="sw">
<polyline points="-48 32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 32 -24 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 80 -32 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 80 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 36 20 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 72 -40 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-44 76 -44 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 40 -40 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="72" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="20" cy="60" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"sw" WINDOW 0 24 16 Left 2
"sw" WINDOW 3 24 96 Left 2
"sw" SYMATTR Value SW
"sw" SYMATTR Prefix S
"sw" SYMATTR Description Voltage controlled switch
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="tline">
<polyline points="-48 16 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 8 32 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-32 -8 32 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -16 48 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 16 -32 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-48 -16 -32 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 16 32 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="48 -16 32 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"tline" WINDOW 0 0 -16 Bottom 2
"tline" WINDOW 3 0 16 Top 2
"tline" SYMATTR Value Td=50n Z0=50
"tline" SYMATTR Prefix T
"tline" SYMATTR Description Ideal Lossless Transmission Line
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="TVSdiode">
<polyline points="0 32 -4 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 32 36 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 32 32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 8 32 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 8 16 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 8 16 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 32 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 56 16 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 56 16 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 56 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"TVSdiode" WINDOW 0 24 -2 Left 2
"TVSdiode" WINDOW 3 24 65 Left 2
"TVSdiode" SYMATTR Value D
"TVSdiode" SYMATTR Prefix D
"TVSdiode" SYMATTR Description Zener Diode
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="varactor">
<polyline points="0 36 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 32 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 16 16 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 16 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"varactor" WINDOW 0 24 0 Left 2
"varactor" WINDOW 3 24 64 Left 2
"varactor" SYMATTR Value D
"varactor" SYMATTR Prefix D
"varactor" SYMATTR Description Diode
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="voltage">
<polyline points="-8 36 8 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="-8 76 8 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 28 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"voltage" WINDOW 0 24 16 Left 2
"voltage" WINDOW 3 24 96 Left 2
"voltage" SYMATTR Value V
"voltage" SYMATTR Prefix V
"voltage" SYMATTR Description Voltage Source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>
<defs>
<g id="zener">
<polyline points="0 44 -4 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 44 36 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="32 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="0 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
<!-- default text positions and size
"zener" WINDOW 0 24 0 Left 2
"zener" WINDOW 3 24 64 Left 2
"zener" SYMATTR Value D
"zener" SYMATTR Prefix D
"zener" SYMATTR Description Zener Diode
-->
<!-- invisible rect added to keep on grid in Inkscape -->
	<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
	</g>
	</defs>

	<defs>
	<g id="and">
	<polyline points="-32 32 -12 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 96 -12 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 96 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 48 32 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="24" cy="80" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -12.00 96.00 A 32.00 32.00 0 0 0 -12.00 32.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"and" WINDOW 0 16 24 Left 2
	"and" WINDOW 3 16 112 Left 2
	"and" SYMATTR Prefix A
	"and" SYMATTR SpiceModel AND
	"and" SYMATTR Description Behavioral AND gate
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="buf">
	<polyline points="0 32 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 48 64 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="64 80 60 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="52" cy="80" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"buf" WINDOW 0 8 16 Left 2
	"buf" WINDOW 3 8 120 Left 2
	"buf" SYMATTR Prefix A
	"buf" SYMATTR SpiceModel BUF
	"buf" SYMATTR Description Behavioral buffer with complementary outputs
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="buf1">
	<polyline points="0 48 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="40 64 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"buf1" WINDOW 0 8 32 Left 2
	"buf1" WINDOW 3 8 104 Left 2
	"buf1" SYMATTR Prefix A
	"buf1" SYMATTR SpiceModel BUF
	"buf1" SYMATTR Description Behavioral buffer
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="counter">
	<polyline points="-80 24 -72 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 40 -72 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 -64 144 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 -64 -80 128" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 128 144 128" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="144 128 144 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"counter" WINDOW 0 -64 -80 Left 2
	"counter" WINDOW 3 -64 152 Left 2
	"counter" SYMATTR Prefix A
	"counter" SYMATTR SpiceModel COUNTER
	"counter" SYMATTR Description Behavioral counter
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="dflop">
	<polyline points="-80 0 80 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 0 -80 144" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 144 80 144" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="80 144 80 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 88 -72 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 104 -72 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="88" cy="96" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"dflop" WINDOW 0 8 -16 Left 2
	"dflop" WINDOW 3 8 168 Left 2
	"dflop" SYMATTR Prefix A
	"dflop" SYMATTR SpiceModel DFLOP
	"dflop" SYMATTR Description Behavioral D-flipflop
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="diffschmitt">
	<polyline points="8 72 16 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 56 20 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 72 12 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 72 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 32 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 48 64 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="64 80 60 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 48 12 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 44 8 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 80 12 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="52" cy="80" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"diffschmitt" WINDOW 0 8 16 Left 2
	"diffschmitt" WINDOW 3 8 120 Left 2
	"diffschmitt" SYMATTR Prefix A
	"diffschmitt" SYMATTR SpiceModel SCHMITT
	"diffschmitt" SYMATTR Description Behavioral Schmitt-Triggered Buffer with Complementary Outs and Differential Input
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="diffschmtbuf">
	<polyline points="8 72 16 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 56 20 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 72 12 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 72 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 32 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 48 12 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 44 8 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 80 12 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"diffschmtbuf" WINDOW 0 8 16 Left 2
	"diffschmtbuf" WINDOW 3 8 120 Left 2
	"diffschmtbuf" SYMATTR Prefix A
	"diffschmtbuf" SYMATTR SpiceModel SCHMITT
	"diffschmtbuf" SYMATTR Description Behavioral Schmitt-Triggered Buffer with Differential Input
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="diffschmtinv">
	<polyline points="8 72 16 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 56 20 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 72 12 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 72 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 32 44 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 44 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="60 64 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 48 12 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 44 8 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 80 12 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="52" cy="64" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"diffschmtinv" WINDOW 0 8 16 Left 2
	"diffschmtinv" WINDOW 3 8 120 Left 2
	"diffschmtinv" SYMATTR Prefix A
	"diffschmtinv" SYMATTR SpiceModel SCHMITT
	"diffschmtinv" SYMATTR Description Behavioral Schmitt-Triggered Inverter with Differential Input
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="inv">
	<polyline points="0 48 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="56 64 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="48" cy="64" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"inv" WINDOW 0 8 32 Left 2
	"inv" WINDOW 3 8 104 Left 2
	"inv" SYMATTR Prefix A
	"inv" SYMATTR SpiceModel BUF
	"inv" SYMATTR Description Behavioral Inverter
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="or">
	<polyline points="-32 32 -28 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 96 -28 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 48 32 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 80 -24 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 64 -20 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 48 -24 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="24" cy="80" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -30.71 96.94 A 56.00 56.00 0 0 0 -30.71 31.06" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -28.00 95.96 A 54.00 54.00 0 0 0 23.43 63.75" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 23.43 64.25 A 54.00 54.00 0 0 0 -28.00 32.04" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"or" WINDOW 0 -8 8 Left 2
	"or" WINDOW 3 -8 128 Left 2
	"or" SYMATTR Prefix A
	"or" SYMATTR SpiceModel OR
	"or" SYMATTR Description Behavioral OR gate with complementary outputs
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="phidet">
	<polyline points="-32 -24 -24 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 -8 -24 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 24 -24 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 8 -24 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 -48 -32 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 -48 32 -48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 48 32 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 48 64 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 -48 64 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 -16 28 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 16 28 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="12" cy="0" rx="24" ry="24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="68" cy="0" rx="12" ry="12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="84" cy="0" rx="12" ry="12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"phidet" WINDOW 0 8 -64 Center 2
	"phidet" WINDOW 3 0 64 Left 2
	"phidet" SYMATTR Prefix A
	"phidet" SYMATTR SpiceModel PHASEDET
	"phidet" SYMATTR Description Behavioral Type 3/4 Phase Detector (phase/frequency detector).  NOTE: Limit input rise times
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="schmitt">
	<polyline points="8 72 16 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 56 20 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 72 12 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 72 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 32 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 48 64 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="64 80 60 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="52" cy="80" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"schmitt" WINDOW 0 8 16 Left 2
	"schmitt" WINDOW 3 8 120 Left 2
	"schmitt" SYMATTR Prefix A
	"schmitt" SYMATTR SpiceModel SCHMITT
	"schmitt" SYMATTR Description Behavioral Schmitt-Triggered buffer with complementary outs
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="schmtbuf">
	<polyline points="4 68 12 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 60 16 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 60 8 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 68 12 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 48 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="40 64 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"schmtbuf" WINDOW 0 8 32 Left 2
	"schmtbuf" WINDOW 3 8 104 Left 2
	"schmtbuf" SYMATTR Prefix A
	"schmtbuf" SYMATTR SpiceModel SCHMITT
	"schmtbuf" SYMATTR Description Behavioral Schmitt-Triggered buffer
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="schmtinv">
	<polyline points="4 68 12 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 60 16 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 60 8 68" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 68 12 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 48 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 40 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="56 64 64 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="48" cy="64" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"schmtinv" WINDOW 0 8 32 Left 2
	"schmtinv" WINDOW 3 8 104 Left 2
	"schmtinv" SYMATTR Prefix A
	"schmtinv" SYMATTR SpiceModel SCHMITT
	"schmtinv" SYMATTR Description Behavioral Schmitt-Triggered inverter
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="srflop">
	<polyline points="-48 16 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 16 -48 128" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 128 48 128" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 128 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="56" cy="96" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"srflop" WINDOW 0 -40 0 Left 2
	"srflop" WINDOW 3 -40 152 Left 2
	"srflop" SYMATTR Prefix A
	"srflop" SYMATTR SpiceModel SRFLOP
	"srflop" SYMATTR Description Behavioral Set-Reset Flipflop
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="xor">
	<polyline points="-36 96 -28 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-36 32 -28 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 48 16 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 80 -40 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 64 -36 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 48 -40 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="24" cy="80" rx="8" ry="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -46.71 96.94 A 56.00 56.00 0 0 0 -46.71 31.06" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -34.71 96.94 A 56.00 56.00 0 0 0 -34.71 31.06" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -28.00 95.96 A 54.00 54.00 0 0 0 23.43 63.75" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 23.43 64.25 A 54.00 54.00 0 0 0 -28.00 32.04" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"xor" WINDOW 0 16 24 Left 2
	"xor" WINDOW 3 16 112 Left 2
	"xor" SYMATTR Prefix A
	"xor" SYMATTR SpiceModel XOR
	"xor" SYMATTR Description Behavioral XOR gate
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="battery">
	<polyline points="-32 36 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 60 32 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 16 0 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 24 -12 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-16 20 -16 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<rect x="-16" y="44" width="32" height="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<rect x="-16" y="68" width="32" height="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"battery" WINDOW 0 24 16 Left 2
	"battery" WINDOW 3 24 104 Left 2
	"battery" SYMATTR Value V
	"battery" SYMATTR Prefix V
	"battery" SYMATTR Description Voltage Source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="cell">
	<polyline points="0 40 0 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 24 32 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 0 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 12 -12 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-16 8 -16 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<rect x="-16" y="32" width="32" height="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"cell" WINDOW 0 24 8 Left 2
	"cell" WINDOW 3 24 56 Left 2
	"cell" SYMATTR Value V
	"cell" SYMATTR Prefix V
	"cell" SYMATTR Description Voltage Source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="DIAC">
	<polyline points="0 44 36 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 20 36 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="36 20 20 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 20 20 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 0 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 44 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 44 64 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 44 44 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="44 20 60 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="36 20 64 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"DIAC" WINDOW 0 48 0 Left 2
	"DIAC" WINDOW 3 48 72 Left 2
	"DIAC" SYMATTR Value DIAC
	"DIAC" SYMATTR Prefix X
	"DIAC" SYMATTR Description Generic Bi-directional Trigger Device symbol for use with a model that you supply
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="Epoly">
	<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 72 4 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 40 4 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 36 0 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"Epoly" WINDOW 0 24 16 Left 2
	"Epoly" WINDOW 3 24 104 Left 2
	"Epoly" SYMATTR Value POLY()
	"Epoly" SYMATTR Prefix E
	"Epoly" SYMATTR Description Voltage dependent voltage source with two terminals.  Useful for drafting schematic generating archaic voltage controlled syntax.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="EuropeanCap">
	<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<rect x="0" y="36" width="32" height="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<rect x="0" y="20" width="32" height="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"EuropeanCap" WINDOW 0 24 4 Left 2
	"EuropeanCap" WINDOW 3 24 64 Left 2
	"EuropeanCap" SYMATTR Value C
	"EuropeanCap" SYMATTR Prefix C
	"EuropeanCap" SYMATTR Description Polarized Capacitor
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="EuropeanPolcap">
	<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 28 32 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 28 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 28 0 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 37 0 37" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 38 0 38" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 39 0 39" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 40 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 41 0 41" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 42 0 42" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 43 0 43" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<rect x="0" y="36" width="32" height="8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"EuropeanPolcap" WINDOW 0 24 4 Left 2
	"EuropeanPolcap" WINDOW 3 24 64 Left 2
	"EuropeanPolcap" SYMATTR Value C
	"EuropeanPolcap" SYMATTR Prefix C
	"EuropeanPolcap" SYMATTR Description Polarized Capacitor
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="EuropeanResistor">
	<polyline points="16 88 16 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 16 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="27 24 5 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="5 88 5 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="5 88 27 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="27 24 27 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"EuropeanResistor" WINDOW 0 31 40 Left 2
	"EuropeanResistor" WINDOW 3 31 76 Left 2
	"EuropeanResistor" SYMATTR Value R
	"EuropeanResistor" SYMATTR Prefix R
	"EuropeanResistor" SYMATTR Description A Resistor(European style graphic)
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="fixedind">
	<path d=" M 4.69 67.31 A 16.00 16.00 0 1 0 4.69 44.69" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 16.00 96.00 A 16.00 16.00 0 1 0 4.69 68.69" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 4.69 43.31 A 16.00 16.00 0 1 0 16.00 16.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"fixedind" WINDOW 0 36 40 Left 2
	"fixedind" WINDOW 3 36 80 Left 2
	"fixedind" SYMATTR Value L
	"fixedind" SYMATTR Prefix L
	"fixedind" SYMATTR Description An inductor that cannot be edited.  Typically used with the 1533/1534
	"fixedind" SYMATTR Def_Sub noedit
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="Gpoly">
	<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 52 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 52 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 52 4 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 52 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"Gpoly" WINDOW 0 24 16 Left 2
	"Gpoly" WINDOW 3 24 104 Left 2
	"Gpoly" SYMATTR Value POLY()
	"Gpoly" SYMATTR Prefix G
	"Gpoly" SYMATTR Description Voltage dependent current source with two terminals.  Useful for drafting schematic generating archaic voltage controlled syntax.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="jumper">
	<polyline points="24 64 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 64 -32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="-20" cy="64" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="20" cy="64" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 19.61 60.08 A 20.00 20.00 0 0 0 -19.61 60.08" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"jumper" WINDOW 0 0 40 Bottom 2
	"jumper" SYMATTR Prefix J
	"jumper" SYMATTR Description A wire jumper.  This component lets you give the same net two different names
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="neonbulb">
	<polyline points="0 -48 0 -21" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 48 0 21" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 8 12 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 12 16 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="19 11 13 5" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="13 11 19 5" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="0" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="-16" rx="5" ry="5" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="16" rx="5" ry="5" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="16" cy="8" rx="1" ry="1" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="16" cy="8" rx="2" ry="2" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="16" cy="8" rx="3" ry="3" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="16" cy="8" rx="4" ry="4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"neonbulb" WINDOW 0 16 -48 Left 2
	"neonbulb" SYMATTR Prefix X
	"neonbulb" SYMATTR Description Parameterized Neon Bulb
	"neonbulb" SYMATTR Value2 Vstrike=100 Vhold=50
	"neonbulb" SYMATTR ModelFile neonbulb.sub
	"neonbulb" SYMATTR SpiceModel neonbulb
	"neonbulb" SYMATTR SpiceLine Zon=2K Ihold=200u
	"neonbulb" SYMATTR SpiceLine2 Tau=100u
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="NIGBT">
	<polyline points="16 80 32 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 96 28 92" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 96 32 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 92 32 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 48 40 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 8 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 40 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 72 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 16 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 0 32 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 16 36 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 16 32 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="36 12 32 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"NIGBT" WINDOW 0 56 32 Left 2
	"NIGBT" WINDOW 3 56 72 Left 2
	"NIGBT" SYMATTR Value NIGBT
	"NIGBT" SYMATTR Prefix Z
	"NIGBT" SYMATTR Description N-Channel IGBT symbol for a model you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="pentode">
	<polyline points="-48 -16 -48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 -16 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 -64 0 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -32 20 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -28 20 -28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -32 -20 -28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 -32 20 -28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 -16 12 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 -16 -4 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-12 -16 -20 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 -16 -48 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 0 28 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 0 12 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 0 -4 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-12 0 -20 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 16 -28 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 16 -12 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 16 4 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 16 20 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 28 24 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 64 -32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 28 -32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 28 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 32 28 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 48.00 -16.00 A 48.00 48.00 0 0 0 -48.00 -16.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -48.00 16.00 A 48.00 48.00 0 0 0 48.00 16.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"pentode" WINDOW 0 8 -80 Left 2
	"pentode" WINDOW 3 -24 80 Left 2
	"pentode" SYMATTR Value Pentrode
	"pentode" SYMATTR Prefix X
	"pentode" SYMATTR Description This symbol is for use with a subcircuit macromodel that you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="PIGBT">
	<polyline points="16 80 36 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 80 32 92" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="36 84 32 92" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 96 32 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 48 40 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 8 16 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 40 16 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 72 16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 80 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 16 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 16 32 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 4 48 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 12 48 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 12 28 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"PIGBT" WINDOW 0 56 32 Left 2
	"PIGBT" WINDOW 3 56 72 Left 2
	"PIGBT" SYMATTR Value PIGBT
	"PIGBT" SYMATTR Prefix Z
	"PIGBT" SYMATTR Description P-Channel IGBT symbol for a model you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="SCR">
	<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 20 16 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 0 16 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-12 64 -32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-12 64 8 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"SCR" WINDOW 0 24 0 Left 2
	"SCR" WINDOW 3 24 72 Left 2
	"SCR" SYMATTR Value SCR
	"SCR" SYMATTR Prefix X
	"SCR" SYMATTR Description Generic SCR symbol for use with a model that you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="signal">
	<polyline points="-4 32 4 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 80 4 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 28 0 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="56" rx="32" ry="32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 0.00 56.00 A 12.00 12.00 0 0 0 24.00 56.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 0.00 56.00 A 12.00 12.00 0 0 0 -24.00 56.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"signal" WINDOW 0 24 16 Left 2
	"signal" WINDOW 3 24 104 Left 2
	"signal" SYMATTR Value SINE(0 1 1K)
	"signal" SYMATTR Prefix V
	"signal" SYMATTR Description Voltage Source, either DC, AC, PULSE, SINE, PWL, EXP, or SFFM
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="tetrode">
	<polyline points="-48 0 -48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 0 48 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 -48 0 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -16 20 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -12 20 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -16 -20 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 -16 20 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="48 0 28 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 0 12 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 0 -4 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-12 0 -20 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 0 -36 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-48 16 -28 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 16 -12 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 16 4 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 16 20 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 16 36 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 28 24 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 64 -32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 28 -32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 28 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 32 28 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M 48.00 0.00 A 48.00 48.00 0 0 0 -48.00 -0.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<path d=" M -48.00 16.00 A 48.00 48.00 0 0 0 48.00 16.00" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"tetrode" WINDOW 0 8 -64 Left 2
	"tetrode" WINDOW 3 -24 80 Left 2
	"tetrode" SYMATTR Value Tetrode
	"tetrode" SYMATTR Prefix X
	"tetrode" SYMATTR Description This symbol is for use with a subcircuit macromodel that you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="TRIAC">
	<polyline points="0 44 36 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 20 36 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="36 20 20 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 20 20 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 0 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 44 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 44 64 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 44 44 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="44 20 60 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="36 20 64 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 64 -16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 64 20 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"TRIAC" WINDOW 0 48 0 Left 2
	"TRIAC" WINDOW 3 48 72 Left 2
	"TRIAC" SYMATTR Value TRIAC
	"TRIAC" SYMATTR Prefix X
	"TRIAC" SYMATTR Description Generic TRIAC symbol for use with a model that you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="triode">
	<polyline points="-48 0 -28 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 0 -12 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 0 4 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="12 0 20 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="28 0 36 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 -48 0 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -16 20 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -12 20 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 -16 -20 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="20 -16 20 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 12 24 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 48 -32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 12 -32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 12 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 16 28 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="0" cy="0" rx="48" ry="48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"triode" WINDOW 0 8 -64 Left 2
	"triode" WINDOW 3 -24 64 Left 2
	"triode" SYMATTR Value Triode
	"triode" SYMATTR Prefix X
	"triode" SYMATTR Description This symbol is for use with a subcircuit macromodel that you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="urc">
	<polyline points="0 72 16 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 56 0 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 40 32 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 24 0 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 16 32 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-8 16 -8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-8 48 -16 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"urc" WINDOW 0 36 32 Left 2
	"urc" WINDOW 3 36 72 Left 2
	"urc" SYMATTR Value URC
	"urc" SYMATTR Prefix U
	"urc" SYMATTR Description Uniform RC-line.  Intended for interconnection on IC's but rarely used.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="urc2">
	<polyline points="16 0 32 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 12 32 4" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 12 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 28 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 28 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 44 32 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 44 32 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 60 32 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 60 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-16 32 -8 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-8 4 -8 60" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"urc2" WINDOW 0 36 16 Left 2
	"urc2" WINDOW 3 36 56 Left 2
	"urc2" SYMATTR Value URC
	"urc2" SYMATTR Prefix U
	"urc2" SYMATTR Description Uniform RC-line.  Intended for interconnection on IC's but rarely used.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="xtal">
	<polyline points="16 44 16 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 20 16 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 20 32 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 44 32 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 36 24 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 36 24 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 28 8 28" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 28 8 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"xtal" WINDOW 0 24 0 Left 2
	"xtal" WINDOW 3 24 72 Left 2
	"xtal" SYMATTR Value C
	"xtal" SYMATTR Prefix C
	"xtal" SYMATTR Description Piezoelectric crystal.  Set C, Lser and Cpar to set series and parallel resonances.  This is actually the same as a capacitor.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="xvaristor">
	<polyline points="16 80 16 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 32 8 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 32 8 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="16 16 16 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 16 0 24" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 24 32 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="32 92 32 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="24 80 24 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 80 24 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"xvaristor" WINDOW 0 36 40 Left 2
	"xvaristor" WINDOW 3 36 76 Left 2
	"xvaristor" SYMATTR Value varistor
	"xvaristor" SYMATTR Prefix X
	"xvaristor" SYMATTR Description Generic varistor symbol for use with a model that you supply.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="UniversalOpamp2">
	<polyline points="-32 -32 32 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 32 32 0" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 -32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 -16 -20 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 16 -20 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 20 -24 12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 -32 0 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 32 0 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 -20 12 -20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 -24 8 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 20 12 20" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"UniversalOpamp2" WINDOW 0 16 -32 Left 2
	"UniversalOpamp2" SYMATTR SpiceModel level.2
	"UniversalOpamp2" SYMATTR Prefix X
	"UniversalOpamp2" SYMATTR Description Universal Opamp model that allows 4 different levels of simulation accuracy.  See ./examples/Educational/UniversalOpamp2.asc for details.  En and in are equivalent voltage and current noises.  Enk and ink are the respective corner frequencies.  Phimargin is used to set the 2nd pole or delay to the approximate phase margin for level.3a and level.3b.  This version uses the new, experimental level 2 switch as the output devices.
	"UniversalOpamp2" SYMATTR Value2 Avol=1Meg GBW=10Meg Slew=10Meg
	"UniversalOpamp2" SYMATTR SpiceLine ilimit=25m rail=0 Vos=0 phimargin=45
	"UniversalOpamp2" SYMATTR SpiceLine2 en=0 enk=0 in=0 ink=0 Rin=500Meg
	"UniversalOpamp2" SYMATTR ModelFile UniversalOpamps2.sub
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="opamp">
	<polyline points="-32 32 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 96 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 32 -32 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 48 -20 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 80 -20 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 84 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"opamp" WINDOW 0 0 32 Left 2
	"opamp" SYMATTR Prefix X
	"opamp" SYMATTR Description Ideal single-pole operational amplifier. You must .lib opamp.sub
	"opamp" SYMATTR Value opamp
	"opamp" SYMATTR SpiceLine Aol=100K
	"opamp" SYMATTR SpiceLine2 GBW=10Meg
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="opamp2">
	<polyline points="-32 32 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 96 32 64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 32 -32 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 48 -20 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 80 -20 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 84 -24 76" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 32 0 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 96 0 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 44 12 44" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 40 8 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 84 12 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"opamp2" WINDOW 0 16 32 Left 2
	"opamp2" WINDOW 3 16 96 Left 2
	"opamp2" SYMATTR Value opamp2
	"opamp2" SYMATTR Prefix X
	"opamp2" SYMATTR Description Basic Operational Amplifier symbol for use with subcircuits in the file ./lib/sub/LTC.lib.  You must give the value a name and include this file.
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="varistor">
	<polyline points="-16 96 -16 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-16 32 -16 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 96 -4 92" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 36 -28 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-4 92 -28 36" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 80 -24 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 48 -24 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 88 -8 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-8 40 -8 88" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-8 40 -24 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 88 -24 40" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-20 80 -12 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-16 76 -16 84" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"varistor" WINDOW 0 0 32 Left 2
	"varistor" WINDOW 3 0 104 Left 2
	"varistor" SYMATTR Prefix A
	"varistor" SYMATTR SpiceModel VARISTOR
	"varistor" SYMATTR Description Voltage controlled varistor
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="sample">
	<polyline points="-80 -64 80 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 -64 -80 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 96 80 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="80 96 96 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="96 16 80 -64" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 40 -72 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-80 24 -72 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"sample" WINDOW 0 -64 -80 Left 2
	"sample" WINDOW 3 -64 120 Left 2
	"sample" SYMATTR Prefix A
	"sample" SYMATTR SpiceModel SAMPLEHOLD
	"sample" SYMATTR Description Behavioral Sample and Hold function block
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="ota2">
	<polyline points="-32 32 12 56" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 96 12 72" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 32 -32 96" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 48 -20 48" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 80 -20 80" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 44 -24 52" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="20" cy="64" rx="12" ry="12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="36" cy="64" rx="12" ry="12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"ota2" WINDOW 0 0 32 Left 2
	"ota2" WINDOW 3 0 104 Left 2
	"ota2" SYMATTR Prefix A
	"ota2" SYMATTR SpiceModel OTA
	"ota2" SYMATTR Description Behavioral Operational Transconductance Amplifier
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	<defs>
	<g id="odev2">
	<polyline points="-32 -32 12 -8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 32 12 8" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-32 -32 -32 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 16 -20 16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-28 -16 -20 -16" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="-24 -20 -24 -12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 -15 0 -32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="0 15 0 32" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 -18 12 -18" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="8 -14 8 -22" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<polyline points="4 18 12 18" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="20" cy="0" rx="12" ry="12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<ellipse cx="36" cy="0" rx="12" ry="12" style="fill:none;fill-opacity:1;stroke:#000000;stroke-width:2.5;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:0;stroke-dasharray:0;stroke-dashoffset:0;stroke-opacity:1" />
	<!-- default text positions and size
	"odev2" WINDOW 0 16 -32 Left 2
	"odev2" WINDOW 3 16 32 Left 2
	"odev2" SYMATTR Prefix O
	"odev2" SYMATTR Description Transconductance Device\n\nParameters:\n  Gm1=<1st order transconductance>\n  Gm2=<2nd order transconductance>\n  Ibias=<Quiescent burn current>
	-->
	<!-- invisible rect added to keep on grid in Inkscape -->
		<rect x="-32" y="-32" width="160" height="160" style="fill:none;stroke:none" />
		</g>
		</defs>
	

`
	return labelDefn
}
