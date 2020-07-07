// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "234b4e270f524d8ef33014c5c4c8ccde"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"15073accf79f3f91b90312cd588f1849": "1f8b08000000000000ff9492416bdb401085cf9a5f31e85024a3aeee011f9ad83439c494a6b752ca489ac88ba5dd6567e4b618fdf722c531760a2e398d34bbefbd6f772750bda396f170301bea791c016c1f7c54cc20a97a4c5babdba132b5efcbd67fdc45522fe5b1845d5b3ad672ab1acaaaa3867b12e5984272aeb3cad1bab694a1929a5cd9fb86bbffec09ddd05a27a5a88fd4720a8888e989728efca7291cf7b6e637d6b2f541c2ecde706d7bea52c801f61451e21e174795797aa900fa27cc17f2e89bbb8e44c61145e3502b1e2069f0886456e421615c54bd59bbd63a8611e07970356ef85796e3e2d2e20049641da2c30f170b87f1a4cbe88d28c7076775453ebb88cd27b3897d89aff05324992687844c834b6ca68f47db4652ce7218af06dcab86ecfc24730019c625f235ba599823c7e8e324995ec57cf58372cc24ee0b24c3f9e9dcce76d7ccbe445fb3c8fab746ebc4d69904aed13a2d905f7bb89887c7dc6dc9bad3ce0279cf4e05bfff385f9e7a6770efa398c433418155e7ebdd37dbb328f5e18568eedd936ca7d1b0ae3d225cf24d9d029f99f1387866f552df4d757ac979042bbc592299c6ac6eb37cfa374fac59dafad8df28551dfff441ad77921698ae379f1f36ebe583737e759be6e6d3a0fed50e921c6084bf010000ffff808732ea05040000",
		"199ad9abb6909544028bc5dd60e13162": "1f8b08000000000000ff5cca3b8ec2301000d07ae71453ee16f127d9e07009eee024a3c122f6087f105294bb53d020aad7bc286bdb08f75d5d7ca4e3006041abec0090e9de4226fc0544440ef5da66b548d42cdd2dfb2a45bfc1e84ba5fcdd42a51c12ebd2e6b2f8f4b17e58369f584966fdd4892a3e8c32ca74bdb16773ea273b8dc3ffd8adde0e6e7193b3ab833f8057000000ffffc1665f1aad000000",
		"19cc0e1aa84ea3978bd983e5c849ace9": "1f8b08000000000000ff2a484cce4e4c4f554849cce7e24a2bcd4b5670c9d7d0ace6e2aa05040000ffffd318594c1a000000",
		"1db18360e80cf6aff78b6551ad6ebb67": "1f8b08000000000000ff548c318ec23010456bff538c52ac92d52aeeb7a6a6e1048e33180b625be3310845b93b05a1a0fdffbd579cbfbac05459eed13310979245a987e942d44b9b469f171b9525a6606b9baa77c9965b0b31555b358b0bdc6100f459984eef0e5595e69556989976683cb88c0d38b7e4e9c88ffeeb19e8f7e3ae30c2da24d1cf3ead3066fea7f90f66c38657000000ffff3122ce6bb5000000",
		"53c9328d2bec1463442936e123699a5a": "1f8b08000000000000ff64cd3d6e03211005e09a39c513d5ae95401fc965942e457202964c30daf0236037c58abb47d8561ad33c896fe64d3676358e71692d13f990536998482c01d2f976d9166553d02e3dafc5b454f53df2ea74e4a6c79e5e7ecc1707531b1749421e877a37817bd795cbee2d4b9a8976534671dd2d4ef77ff579cbc1df5bb4f8485be332d5878927304e4b50afd1f9c8338e5bd1199584c3cb19acde4adaf224b5c95ece240e12a20e71ff52ad8983ae06005ae3b7f8c628d7bb50e391109d44a74e7f010000ffffd6794f5e1c010000",
		"7e573a542eb21fa4bbc259a7b00dbcca": "1f8b08000000000000ff2a484cce4e4c4f55c8cd4f49cde1e22aa92c4855a8aed6f3cd4f71ce492c2eaead55282e292a4d2e51a8e6aae502040000ffffcb815bad2d000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("all", "./templates/plugin")
		b.SetResolver("dao/dao.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "19cc0e1aa84ea3978bd983e5c849ace9"})
		b.SetResolver("go.mod.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "199ad9abb6909544028bc5dd60e13162"})
		b.SetResolver("http/http.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "53c9328d2bec1463442936e123699a5a"})
		b.SetResolver("main.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "15073accf79f3f91b90312cd588f1849"})
		b.SetResolver("model/model.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "7e573a542eb21fa4bbc259a7b00dbcca"})
		b.SetResolver("service/service.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "1db18360e80cf6aff78b6551ad6ebb67"})
	}()

	return nil
}()