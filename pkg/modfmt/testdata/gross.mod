

// This file is a complete mess :)
// This is an example header comment.

  //   Here is another header comment.

// module comment 1
module example.com/my/module //   module comment 2


		//   And another header comment.

		//   go comment 1
go 1.23.2 // go comment 2

		// toolchain comment 1
toolchain go1.23 // toolchain comment 2

// godebug foo comment 1
godebug foo=true    //   godebug foo comment 2

		godebug (

		// godebug bar comment 1
		bar=true   // godebug bar comment 2
		)

// replace example.com/a/a comment 1
replace example.com/a/a v1.0.0 => example.com/c/c v1.0.0-alpha.1  // replace example.com/a/a comment 2

		tool (
		// tool example.com/tool/b comment 1
		example.com/tool/b   // tool example.com/tool/b comment 2
		// tool example.com/tool/a comment 1
		//   tool example.com/tool/a comment 2
		example.com/tool/a
		)


		// replace local example.com/d/d comment 1
	replace  	example.com/d/d => ./local/d   // replace local example.com/d/d comment 2

// This comment will be dropped
require (
// require indirect example.com/a/a comment 1
		example.com/a/a v1.1.1 // indirect
		// require example.com/b/b comment 1
example.com/b/b v1.2.2 // require example.com/b/b comment 1

		// require indirect example.com/b/b comment 1
		example.com/b/b v1.2.2 // indirect
		)

	// This comment will also be dropped
replace (
		// replace local example.com/b/b comment 1
  // replace local example.com/b/b comment 2
		example.com/b/b => ./local/b

		// replace example.com/a/a comment 1
		example.com/a/a => example.com/b/b v1.0.0-alpha.1  // replace example.com/a/a comment 2
)


ignore (
// ignore b comment 1
		// ignore b comment 2
		./ignore/b

// ignore a comment 1
./ignore/a   // ignore b comment 2
)


		// require example.com/a/a comment 1
require 		example.com/a/a v1.1.1 // require example.com/a/a comment 2


		exclude (
		// exclude example.com/c/c comment 1
			//    exclude example.com/c/c comment 2
		example.com/c/c v1.3.3
)

		retract (
//retract v1.1.1 comment 1
		v1.1.1 //  retract v1.1.1 comment 2

		// retract v1.20.2 comment 1
		v1.20.2 //  retract v1.20.2 comment 2

		// retract v1.3.0-v1.3.9 comment 1
		[v1.3.0   ,v1.3.9   ] //   retract v1.3.0-v1.3.9 comment 2
		)

// And the final header comment.
