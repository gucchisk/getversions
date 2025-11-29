/*
Copyright © 2023 gucchisk
*/
package cmd

import (
	"github.com/gucchisk/getversions/actions/apache"
)

// var log logr.Logger

// var htmltxt = strings.NewReader(`
// <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
// <html>
//  <head>
//   <title>Index of /maven/maven-3</title>
//  </head>
//  <body>
// <h1>Index of /maven/maven-3</h1>
// <pre><img src="/icons/blank.gif" alt="Icon "> <a href="?C=N;O=D">Name</a>                    <a href="?C=M;O=A">Last modified</a>      <a href="?C=S;O=A">Size</a>  <a href="?C=D;O=A">Description</a><hr><img src="/icons/back.gif" alt="[PARENTDIR]"> <a href="/maven/">Parent Directory</a>                             -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.0.5/">3.0.5/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.1.1/">3.1.1/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.2.5/">3.2.5/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.3.9/">3.3.9/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.5.4/">3.5.4/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.6.3/">3.6.3/</a>                  2022-06-17 11:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.8.8/">3.8.8/</a>                  2023-03-14 11:46    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.9.0/">3.9.0/</a>                  2023-02-06 08:16    -
// <img src="/icons/folder.gif" alt="[DIR]"> <a href="3.9.1/">3.9.1/</a>                  2023-03-18 09:52    -
// <hr></pre>
// </body></html>
// `)

var apacheAction = apache.ApacheAction{}
var apacheCmd = createActionCmd("apache", &apacheAction)
