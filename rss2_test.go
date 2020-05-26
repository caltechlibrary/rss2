//
// rss2 is a golang package for working with RSS 2 feeds and documents.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2018, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package rss2

import (
	"net/url"
	"strings"
	"testing"
)

func TestRSS2(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="utf-8" ?>
<rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>

    <title>CaltechAUTHORS: Title matches "Molecules in solution". Results ordered -Date Deposited. </title>
    <link>http://authors.library.caltech.edu/</link>
    <atom:link xmlns:atom="http://www.w3.org/2005/Atom" rel="self" href="http://authors.library.caltech.edu/cgi/search/advanced/?output=RSS2&amp;title=Molecules+in+solution" type="application/rss+xml"></atom:link>
    <description>1. This is an institutional repository.
2. CaltechAUTHORS holds all types of materials.
3. Deposited items may include:
   (a) working drafts
   (b) submitted versions (as sent to journals for peer-review)
   (c) accepted versions (author's final peer-reviewed drafts)
   (d) published versions (publisher-created files)
4. Items are individually tagged with:
   (a) their version type and date.
   (b) their peer-review status.
   (c) their publication status.
5. Principal Languages: English
</description><image>
        <url>http://authors.library.caltech.edu/images/codalogo.jpg</url>
        <title>CaltechAUTHORS: Title matches "Molecules in solution". Results ordered -Date Deposited. </title>
        <link>http://authors.library.caltech.edu/</link></image>
    <pubDate>Fri, 12 Aug 2016 15:00:00 -0700</pubDate>
    <lastBuildDate>Fri, 12 Aug 2016 15:00:00 -0700</lastBuildDate>
    <language>en</language>
    <copyright></copyright>
<item>
  <pubDate>Mon, 25 Jul 2016 20:48:03 -0700</pubDate>
  <title> Flow-through Capture and in Situ Amplification Can Enable Rapid Detection of a Few Single Molecules of Nucleic Acids from Several Milliliters of Solution </title>
  <link>http://authors.library.caltech.edu/69188/</link>
  <guid>http://authors.library.caltech.edu/69188/</guid>
  <description>  Schlappi, Travis S. and McCalla, Stephanie E. and Schoepp, Nathan G. and Ismagilov, Rustem F.  (2016)  Flow-through Capture and in Situ Amplification Can Enable Rapid Detection of a Few Single Molecules of Nucleic Acids from Several Milliliters of Solution.  Analytical Chemistry .    ISSN 0003-2700.      (In Press)  http://resolver.caltech.edu/CaltechAUTHORS:20160725-102649276 &lt;http://resolver.caltech.edu/CaltechAUTHORS:20160725-102649276&gt;  </description></item>
<item>
  <pubDate>Tue, 05 Aug 2014 15:26:26 -0700</pubDate>
  <title> Note on Dipole Moments of Molecules in Solution </title>
  <link>http://authors.library.caltech.edu/47953/</link>
  <guid>http://authors.library.caltech.edu/47953/</guid>
  <description>  Bauer, S. H.  (1936)  Note on Dipole Moments of Molecules in Solution.  Journal of Chemical Physics, 4  (7).   pp. 458-459.  ISSN 0021-9606.       http://resolver.caltech.edu/CaltechAUTHORS:20140804-165648676 &lt;http://resolver.caltech.edu/CaltechAUTHORS:20140804-165648676&gt;  </description></item>
<item>
  <pubDate>Tue, 07 Aug 2012 17:07:28 -0700</pubDate>
  <title> Solution, surface, and single molecule platforms for the study of DNA-mediated charge transport </title>
  <link>http://authors.library.caltech.edu/32968/</link>
  <guid>http://authors.library.caltech.edu/32968/</guid>
  <description>  Muren, Natalie B. and Olmon, Eric D. and Barton, Jacqueline K.  (2012)  Solution, surface, and single molecule platforms for the study of DNA-mediated charge transport.  Physical Chemistry Chemical Physics, 14  (40).   pp. 13754-13771.  ISSN 1463-9076.  PMCID PMC3478128.      http://resolver.caltech.edu/CaltechAUTHORS:20120807-093450882 &lt;http://resolver.caltech.edu/CaltechAUTHORS:20120807-093450882&gt;  </description></item>
<item>
  <pubDate>Wed, 09 Sep 2009 18:17:34 -0700</pubDate>
  <title> Direct Emission of I_2 Molecule and IO Radical from the Heterogeneous Reactions of Gaseous Ozone with Aqueous Potassium Iodide Solution </title>
  <link>http://authors.library.caltech.edu/15526/</link>
  <guid>http://authors.library.caltech.edu/15526/</guid>
  <description>Sakamoto, Yosuke and Yabushita, Akihiro and Kawasaki, Masahiro and Enami, Shinichi  (2009)  Direct Emission of I_2 Molecule and IO Radical from the Heterogeneous Reactions of Gaseous Ozone with Aqueous Potassium Iodide Solution.  Journal of Physical Chemistry A, 113  (27).   pp. 7707-7713.  ISSN 1089-5639.       http://resolver.caltech.edu/CaltechAUTHORS:20090901-131930555 &lt;http://resolver.caltech.edu/CaltechAUTHORS:20090901-131930555&gt;  </description><media:content url="http://authors.library.caltech.edu/15526/4/preview.png" type="image/png"/></item>
<item>
  <pubDate>Fri, 29 Aug 2008 05:05:45 -0700</pubDate>
  <title> Unimolecular reaction rates in solution and in the isolated molecule: Comparison of diphenyl butadiene nonradiative decay in solutions and supersonic jets </title>
  <link>http://authors.library.caltech.edu/11478/</link>
  <guid>http://authors.library.caltech.edu/11478/</guid>
  <description>  Courtney, S. H. and Fleming, G. R. and Khundkar, L. R. and Zewail, A. H.  (1984)  Unimolecular reaction rates in solution and in the isolated molecule: Comparison of diphenyl butadiene nonradiative decay in solutions and supersonic jets.  Journal of Chemical Physics, 80  (9).   pp. 4559-4560.  ISSN 0021-9606.       http://resolver.caltech.edu/CaltechAUTHORS:COUjcp84 &lt;http://resolver.caltech.edu/CaltechAUTHORS:COUjcp84&gt;  </description><media:content url="http://authors.library.caltech.edu/11478/2/preview.png" type="image/png"/></item>
<item>
  <title> The osmotic pressure of the ions and of the undissociated molecules of salts in aqueous solution </title>
  <link>http://authors.library.caltech.edu/3382/</link>
  <guid>http://authors.library.caltech.edu/3382/</guid>
  <description>  Bates, Stuart J.  (1915)  The osmotic pressure of the ions and of the undissociated molecules of salts in aqueous solution.  Proceedings of the National Academy of Sciences of the United States of America, 1  (6).   pp. 363-368.  ISSN 0027-8424. http://resolver.caltech.edu/CaltechAUTHORS:BATpnas15 &lt;http://resolver.caltech.edu/CaltechAUTHORS:BATpnas15&gt;  </description><media:content url="http://authors.library.caltech.edu/3382/2/preview.png" type="image/png"/></item>
    </channel>
</rss>`)

	r, err := Parse(src)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	results, err := r.Filter([]string{".item[].title"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(results[".item[].title"].([]string)) != len(r.ItemList) {
		t.Errorf("Expected 6 .item[].title, got %s", strings.Join(results[".item[].title"].([]string), "\t"))
		t.FailNow()
	}
	results, err = r.Filter([]string{".item[].link"})
	if err != nil {
		t.Errorf("Expected 6 .item[].link, got %+v", strings.Join(results[".item[].title"].([]string), "\t"))
		t.FailNow()
	}
	for _, link := range results[".item[].link"].([]string) {
		_, err := url.Parse(link)
		if err != nil {
			t.Errorf("expected to parse link %q into url, %s", link, err)
		}
	}
}
