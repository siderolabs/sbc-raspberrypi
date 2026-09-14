## [sbc-raspberrypi 0.2.2](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.2.2) (2026-09-14)

Welcome to the v0.2.2 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Licia Seiker

### Changes
<details><summary>1 commit</summary>
<p>

* [`b1fe7bd`](https://github.com/siderolabs/sbc-raspberrypi/commit/b1fe7bda481a0f1a0e7fc05cf61d93e1b838e975) fix: restore CM5 Ethernet initialization
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.2.1](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.2.1)

## [sbc-raspberrypi 0.2.1](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.2.1) (2026-08-11)

Welcome to the v0.2.1 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Andrey Smirnov
* Noel Georgi
* Maja Bojarska
* Mateusz Urbanek
* Dmitrii Sharshakov
* Christopher Barnes
* Lukasz Raczylo
* Rokoucha
* Andras BALI
* Andreas Freund
* Andreas Lüdeke
* Ansgar Dahlen
* Calin
* Camillo Rossi
* Christian A. Jacobsen
* Dario Emerson
* David Donchez
* Fritz Schaal
* Gregor Gruener
* Jaakko Sirén
* Jean-Francois Roy
* Joakim Nohlgård
* Justin Garrison
* Kai Zhang
* Max Makarov
* Michal Baumgartner
* Nico Berlee
* Serge van Ginderachter
* YANG JOO WOONG
* Zadkiel AHARONIAN
* appkins
* arita
* ctr49
* eseiker
* imusmanmalik
* lmacka
* pranav767

### Changes
<details><summary>4 commits</summary>
<p>

* [`6e3d8d8`](https://github.com/siderolabs/sbc-raspberrypi/commit/6e3d8d8d39133830f13d6d5dd3017c814eb60cd6) fix(docs): broken link
* [`cd1a270`](https://github.com/siderolabs/sbc-raspberrypi/commit/cd1a2702c018125617b4f963460055584fe39fbc) feat: bump dependencies
* [`3efdf1a`](https://github.com/siderolabs/sbc-raspberrypi/commit/3efdf1a043fa9ed07546d7a1225f3539f675653d) fix: generate Pi 5 DTBs from a kernel matching the one Talos ships
* [`967dbb5`](https://github.com/siderolabs/sbc-raspberrypi/commit/967dbb58cc0f0aebf5b8d2a690c827cac868b688) chore: rekres and bump pkg/tools
</p>
</details>

### Changes from siderolabs/pkgs
<details><summary>206 commits</summary>
<p>

* [`7d8b87b`](https://github.com/siderolabs/pkgs/commit/7d8b87bc009b083ab3b5acc5af314069bbf17cce) chore: dependency updates 2026-08-10
* [`0cc868b`](https://github.com/siderolabs/pkgs/commit/0cc868bfcbddf8c1cd9c1edf21fc691a901d917c) feat(kernel): enable CONFIG_FS_ENCRYPTION
* [`e5b0a80`](https://github.com/siderolabs/pkgs/commit/e5b0a808aa90c42aeba21c58f4d9156225fccc09) feat: bump kernel to 6.18.44
* [`6f77a30`](https://github.com/siderolabs/pkgs/commit/6f77a30685770643da6e30d4664efcc8ca47d4e5) feat: enable CONFIG_FSCACHE and CONFIG_NFS_FSCACHE
* [`b0a7e99`](https://github.com/siderolabs/pkgs/commit/b0a7e996f74e8e49d3a653870310c040eb1abe5b) feat: bump kernel to 6.18.43
* [`f489e70`](https://github.com/siderolabs/pkgs/commit/f489e70132e5630f53218a6d4b5fcd98bf32516d) feat: update backportable dependencies
* [`4c14c73`](https://github.com/siderolabs/pkgs/commit/4c14c7380254c92f2fb3523150a46a07687b9a97) feat: support confiuring gnu mirror urls
* [`445e180`](https://github.com/siderolabs/pkgs/commit/445e180c8251c6538992585d684508decde86294) fix: iptables with Cilium
* [`b256af1`](https://github.com/siderolabs/pkgs/commit/b256af14698ee12c3b741a4f00db4c98834884f5) chore: update kernel
* [`1181429`](https://github.com/siderolabs/pkgs/commit/1181429b162e550fa01dfd930d7b7aab2c7bcb95) fix: backport macb IEEE 802.3az EEE support for Raspberry Pi 5 RP1
* [`aa8d7d6`](https://github.com/siderolabs/pkgs/commit/aa8d7d671e103a5393654bb1fd0ef4e4b7ca43b7) feat: bump kernel to 6.18.41
* [`b253464`](https://github.com/siderolabs/pkgs/commit/b253464ac13afa3c67221cd3a6b0883f4a59565a) feat: enable PCF8523 RTC support for arm64
* [`0383b09`](https://github.com/siderolabs/pkgs/commit/0383b0949de84379eb525382fa2c647c31065607) feat: bump kernel to 6.18.40
* [`37184f0`](https://github.com/siderolabs/pkgs/commit/37184f0b0af0cee59305a9607594c23c3a6fead1) feat: enable CONFIG_NFT_SOCKET in the kernel
* [`f78e3dc`](https://github.com/siderolabs/pkgs/commit/f78e3dc66a2085187c2b7763443619793ca37083) feat: update cryptsetup to 2.8.7
* [`8c5831c`](https://github.com/siderolabs/pkgs/commit/8c5831c6d375c3d364bc737eba42dccbd6ce6cd2) feat: update dependencies
* [`0de2a61`](https://github.com/siderolabs/pkgs/commit/0de2a615b37de7e9139cbd49c76b4b57c5da8ecc) feat: update Linux to 6.18.39
* [`c652074`](https://github.com/siderolabs/pkgs/commit/c6520749fb13d2afcaaea6e6b4b37eba9bfa97ce) feat: enable UFSHC and some other options
* [`c4b550c`](https://github.com/siderolabs/pkgs/commit/c4b550c6e33e814874d13903af9b89675ced9ed8) feat: bump dependencies
* [`de07964`](https://github.com/siderolabs/pkgs/commit/de0796462af2effbbf63281a029af216a4bb199d) feat: enable CONFIG_IOMMUFD and CONFIG_VFIO_DEVICE_CDEV
* [`6a9c40c`](https://github.com/siderolabs/pkgs/commit/6a9c40c85c264e0a9394793544ed5af182cddab7) chore: bump tools 2026-07-14
* [`61e3ed9`](https://github.com/siderolabs/pkgs/commit/61e3ed9875a921b4bd073f00fd016329d182b459) chore: bump bldr to v0.6.1
* [`981029e`](https://github.com/siderolabs/pkgs/commit/981029e1c6310e13a886c5094571e2ea3bcfc1cc) chore: bump nvidia to 580.167.08
* [`ff9a355`](https://github.com/siderolabs/pkgs/commit/ff9a3553c755e9d6ce07b60b750f5e425acc99ff) chore: bump deps (minor)
* [`0b2474e`](https://github.com/siderolabs/pkgs/commit/0b2474e2d9a90d0091d733a884c8ca51353dba6b) chore: bump deps (patch)
* [`94a16ef`](https://github.com/siderolabs/pkgs/commit/94a16efc7a04c291ca79e9a01c04d032d5f622ac) chore: bump toolchain
* [`6dcf355`](https://github.com/siderolabs/pkgs/commit/6dcf35514d92d09747007d9038799ae94a50a4bf) feat: enable devmapper plugin in containerd build
* [`03534a5`](https://github.com/siderolabs/pkgs/commit/03534a5182c898bc14b7ff08ff884ef55db3e246) fix: enable CONFIG_IFB as a module
* [`a264237`](https://github.com/siderolabs/pkgs/commit/a2642379f444fa37f50fd125f0c9771a2d2915fb) feat: update DRBD to 9.3.3
* [`3f09c0c`](https://github.com/siderolabs/pkgs/commit/3f09c0c8d9d5522d0f78a86df98971ae7913afea) chore: update toolchain and tools
* [`6c08c46`](https://github.com/siderolabs/pkgs/commit/6c08c469d7124a49160de1a19ca48afa03ed76ff) feat: bump kernel to 6.18.38
* [`7c4fe92`](https://github.com/siderolabs/pkgs/commit/7c4fe92d0d4ae575031cb5757cc89b8b101ba0b3) feat: build runs with libpathrs (only amd64)
* [`8922b6d`](https://github.com/siderolabs/pkgs/commit/8922b6d46e84025c4d3e93bb5e568632b9682d7d) fix: correct finalize destination path
* [`ae0d701`](https://github.com/siderolabs/pkgs/commit/ae0d701b5389244e7230ebc982d777418da643ad) fix: use non-conflicting name for mdadm
* [`5186a65`](https://github.com/siderolabs/pkgs/commit/5186a654815e10f9ebb57dc29305e094fb4b1d8b) feat: add mdadm package for software RAID
* [`e09f9fb`](https://github.com/siderolabs/pkgs/commit/e09f9fb84419401f98550feba0477334cfa6df0a) feat: bump kernel 6.18.37
* [`8d23631`](https://github.com/siderolabs/pkgs/commit/8d23631f1ad2ee444363f8f4f17d4b5691d81dd8) feat: update Linux firmware to 20260622
* [`0343557`](https://github.com/siderolabs/pkgs/commit/03435572bf736a184f32fcdc12aa405b44ba34c1) feat: update runc to 1.5.0
* [`55d3676`](https://github.com/siderolabs/pkgs/commit/55d3676937edd05bc4a470cc83f61dd59db6978c) feat: bump dependencies
* [`ea48e8b`](https://github.com/siderolabs/pkgs/commit/ea48e8b66153a7909eef240bec10949783b6ac87) fix: patch Linux kernel for tunnel metadata buffer overflow
* [`ff80d88`](https://github.com/siderolabs/pkgs/commit/ff80d88660cde1e201ca9bd039e058bbf87f9a4f) feat: add support for AMD XGBE driver
* [`9f8ab22`](https://github.com/siderolabs/pkgs/commit/9f8ab22a224f16d30adba8a0b4c980866866abd0) feat: enable NF_TABLES_ARP option
* [`bedfbeb`](https://github.com/siderolabs/pkgs/commit/bedfbeb545665f6fe8678aa0729dae49f9c1815e) feat: update Linux to 6.18.36
* [`a9f2bb3`](https://github.com/siderolabs/pkgs/commit/a9f2bb37ad8c8958276e3da8bffa366001042227) chore: bump containerd to 2.3.2
* [`73e76f8`](https://github.com/siderolabs/pkgs/commit/73e76f8ec71d93a963d80710706b2138de2c6d74) chore: upgrade runc to 1.5.0-rc.3
* [`28db1ca`](https://github.com/siderolabs/pkgs/commit/28db1ca7b99651482597c0a69445be47f61bca38) chore: update nvidia driver lts to 580.167.08
* [`5df1a44`](https://github.com/siderolabs/pkgs/commit/5df1a442eb041df2ae4e484798d5c28e518f8b36) chore: update zfs to 2.4.3
* [`cd77c4f`](https://github.com/siderolabs/pkgs/commit/cd77c4f0dfd611a67afd2656fde11c5bf197a004) chore: update dependencies 2026-06-16
* [`0f27ecc`](https://github.com/siderolabs/pkgs/commit/0f27ecc3f22c2c526037be262a2678a539e29708) feat: bump runc
* [`d213ff5`](https://github.com/siderolabs/pkgs/commit/d213ff5e7d88436fd4082fb4defdf4c17c23a8ac) feat: bump OpenSSL to 3.6.3
* [`cb713ae`](https://github.com/siderolabs/pkgs/commit/cb713ae48611f7c8ae31a08232cda482f2d6bfc1) feat: bump kernel to 6.18.35
* [`09cb04e`](https://github.com/siderolabs/pkgs/commit/09cb04e048191a92de49261be6291595eda0ffda) fix: avoid page_table_check BUG on time namespace VVAR page
* [`bfb88f6`](https://github.com/siderolabs/pkgs/commit/bfb88f6eac299390bee7a112f01803b1d27fe7e5) feat: add nvidia-fs kernel module
* [`f2850d1`](https://github.com/siderolabs/pkgs/commit/f2850d169be4a440103ac30a8cd82be6ee05110b) feat: enable USB hiddev for apcupsd support
* [`55aa64f`](https://github.com/siderolabs/pkgs/commit/55aa64fa75a804ece0d356d913382b03e198af5e) feat: bump go to 1.26.4
* [`f27dbe1`](https://github.com/siderolabs/pkgs/commit/f27dbe1f053adf9eb1def340fe5fd801f514daa9) feat: bump kernel to 6.18.34
* [`aa9fe00`](https://github.com/siderolabs/pkgs/commit/aa9fe0089646c4917c5968e7b99d2609fd7e7fc3) feat: add DVB USB Modules
* [`0870a4b`](https://github.com/siderolabs/pkgs/commit/0870a4b6b1d9a01bbdf8ed4586465d1caf013a20) feat: bump dependencies
* [`f9134e5`](https://github.com/siderolabs/pkgs/commit/f9134e53a1cb2eb9555753e7206c7cb80178dd7c) fix: enable CONFIG_BCM2712_MIP as built-in in arm64 kernel config
* [`285c6ae`](https://github.com/siderolabs/pkgs/commit/285c6ae7cca6a6f2f3d44f0c7362f0991b4ba9b5) fix: set usermode static helper to machine
* [`bd2a754`](https://github.com/siderolabs/pkgs/commit/bd2a754171fff6eac35fa9de1bf02eb1df16fac1) feat: pre-generate drbd patches using spatch out of tree
* [`898844e`](https://github.com/siderolabs/pkgs/commit/898844e7da04836dfeda8cc7ca7335bef3ca415a) feat: update Linux to 6.18.33
* [`a8dfbf7`](https://github.com/siderolabs/pkgs/commit/a8dfbf7c0e8c490a4c0cd461ebf251af6654fdd0) fix: disable kernel modprobe path
* [`c542950`](https://github.com/siderolabs/pkgs/commit/c542950c046bfd7d90a34c708a00d955db8c2c9a) fix: pull in tools with zstd sbom
* [`c0ec8f3`](https://github.com/siderolabs/pkgs/commit/c0ec8f389155de453b92a6184bea2dccc831573a) feat: enable PPP and INFINIBAND_BNXT_RE
* [`c62c4e1`](https://github.com/siderolabs/pkgs/commit/c62c4e1fb164b02ad941721c6dd50c9706e1156e) feat: update containerd to 2.3.1
* [`270f9f8`](https://github.com/siderolabs/pkgs/commit/270f9f821d1f7d18b4cc229f4c8e0e7f759bd96b) chore: update deps
* [`4f7feb4`](https://github.com/siderolabs/pkgs/commit/4f7feb4103eab8ab0d6eeb04e34a9dc3cf07067c) feat: enable more options for CRI-U checkpoint/restore
* [`87994f7`](https://github.com/siderolabs/pkgs/commit/87994f7fce1012efe898688d7dcc9a80c3637b1b) feat: move autoloadable stuff as modules
* [`80c27f3`](https://github.com/siderolabs/pkgs/commit/80c27f3295b832e44f897681a0eb0218adfe88cb) fix: drop legacy network protocols
* [`fbb7360`](https://github.com/siderolabs/pkgs/commit/fbb73601696a73ed0de65ef95a1c182f65d7afc9) feat: drop legacy iptables/ebtables support
* [`eac5f86`](https://github.com/siderolabs/pkgs/commit/eac5f865f69f8fbedbf5ade03adb17d2c4200c89) feat: bump kernel 6.18.32
* [`d616f6c`](https://github.com/siderolabs/pkgs/commit/d616f6cc9260e3b167e98c66e13f7f39f2fb64c8) feat: update Linux to 6.18.31
* [`02bcfce`](https://github.com/siderolabs/pkgs/commit/02bcfced0848d0d4c066aed993171f40415e1d72) fix: macb silent TX stall on BCM2712/RP1 (v2 patches from netdev)
* [`12ca698`](https://github.com/siderolabs/pkgs/commit/12ca69857d8a71788708753a8da26e8e2ab11983) feat: update ZFS & NVIDIA LTS
* [`9fff943`](https://github.com/siderolabs/pkgs/commit/9fff9435215dc6eb6b96fea566b3aea7f9bda2a7) feat: update Linux to 6.18.30
* [`c5a1685`](https://github.com/siderolabs/pkgs/commit/c5a168538eac21a4b4cfab4cd2c08eb8fdcd05af) feat: move HWMON as modules
* [`b2a45fb`](https://github.com/siderolabs/pkgs/commit/b2a45fb2569aa7cf4b48700d8edf14a4a0728612) feat: move CONFIG_INTEL_IOATDMA as a module
* [`ea8d35f`](https://github.com/siderolabs/pkgs/commit/ea8d35f8a099027960a91dba1a0bbb90415575b9) feat: move ACPI device drivers as modules
* [`501ba58`](https://github.com/siderolabs/pkgs/commit/501ba580184ae4581503628a591e4fc6181e9dd0) feat: move HID quirks as modules
* [`b35312c`](https://github.com/siderolabs/pkgs/commit/b35312c50f460b06d5eaab6c80d2af8fa469de9e) feat: move PS/2 mouse drivers as modules
* [`3a5d9d7`](https://github.com/siderolabs/pkgs/commit/3a5d9d79db30c7659e90b1a35616ec677b766e65) feat: move IPMI driver to be a module
* [`792a69a`](https://github.com/siderolabs/pkgs/commit/792a69a0fbc3f86ea64340eb511e139e2db793a0) feat: disable AGP drivers
* [`99990b4`](https://github.com/siderolabs/pkgs/commit/99990b4b572dcdde7aac145f7a761dce04829eae) feat: move Hyper-V drivers as modules
* [`fb697d6`](https://github.com/siderolabs/pkgs/commit/fb697d66ae2cf7d7c4be8b2d85c3ef5fa90267e4) feat: move Xen frontend drivers as modules
* [`1df1713`](https://github.com/siderolabs/pkgs/commit/1df171313352ca866e8a8c43ee19974897c46630) feat: move ATA / MMC controllers as modules
* [`f7f9341`](https://github.com/siderolabs/pkgs/commit/f7f93412d618d0790e23d58b7ba2d9fd982f320b) feat: move USB class drivers as modules
* [`ba873e9`](https://github.com/siderolabs/pkgs/commit/ba873e9c4716d4e55199f5903b02bf407fdb52aa) feat: move USB host controllers as modules
* [`8f25baa`](https://github.com/siderolabs/pkgs/commit/8f25baadeed5447c2246f4cd1253f39837903c79) feat: move virtio bus stuff as modules
* [`d0c5480`](https://github.com/siderolabs/pkgs/commit/d0c548047d258834d43c64159edb3d079cbefea4) feat: bump kernel to 6.18.29
* [`dfb09f0`](https://github.com/siderolabs/pkgs/commit/dfb09f02d0b87fb2361331733380e64653396e24) feat: bump kernel 6.18.28
* [`c97bc24`](https://github.com/siderolabs/pkgs/commit/c97bc24ea16ab8a66b9ea0478a41b1e421cadfd1) feat: update Go to 1.26.3
* [`dfe8926`](https://github.com/siderolabs/pkgs/commit/dfe8926a029f6360b27afbd0bb0367358b847210) feat: add btrfsprogs
* [`06ff9dc`](https://github.com/siderolabs/pkgs/commit/06ff9dcde8e75e7ccbc7f4b705d4ec220fed0e2d) feat: update Linux to 6.18.27
* [`2265fc9`](https://github.com/siderolabs/pkgs/commit/2265fc968f5dcf30fabe7270930f8b6dece744e3) feat(kernel): backport two PCI bridge realloc fixes from v6.19
* [`5a21d99`](https://github.com/siderolabs/pkgs/commit/5a21d99a6d47a27f3b0a250a76a830e929a608de) feat: bump dependencies
* [`cb3f406`](https://github.com/siderolabs/pkgs/commit/cb3f4069708650a9e4add8f5203d29fbb30c7ca7) feat: update containerd to 2.3.0
* [`e192574`](https://github.com/siderolabs/pkgs/commit/e192574207d1093e9a43a723309b30d028c31b8c) feat: update Linux to 6.18.26
* [`e5e6cb8`](https://github.com/siderolabs/pkgs/commit/e5e6cb816acb97ddca2b1f0ed64897eac0a3eb07) feat: update DRBD to 9.3.2
* [`77538b1`](https://github.com/siderolabs/pkgs/commit/77538b1543a3c2c6ffadd587c47930b96c3b8517) feat: update NVIDIA drivers
* [`adeaafc`](https://github.com/siderolabs/pkgs/commit/adeaafc45f6a87259b51d452a44d279a1d48b0da) feat: preserve System.map on kernel builds
* [`c77f985`](https://github.com/siderolabs/pkgs/commit/c77f9851d2c9e4ac90639d392eeaf27d8319fb0b) fix: disable legacy framebuffer drivers
* [`8f3ef77`](https://github.com/siderolabs/pkgs/commit/8f3ef7751868f62e153f5a679702b3f2c4634c30) fix: enable safesetid LSM
* [`f82d3af`](https://github.com/siderolabs/pkgs/commit/f82d3afec43ef086b1b06b163b94d119e88fc3a3) fix: disable CONFIG_DEVPORT
* [`b189a96`](https://github.com/siderolabs/pkgs/commit/b189a96319de60a6d3cd2261ee7d7c3dad6f3657) fix: disable crypto user API
* [`9a718f6`](https://github.com/siderolabs/pkgs/commit/9a718f6a64aaeb260a9e5182c93817676beff270) docs: list net macb silent TX stall fixes in kernel/build/patches/README.md
* [`ca3599f`](https://github.com/siderolabs/pkgs/commit/ca3599f4b801dfe0218ae0486bbcc0e26761b103) fix: macb silent TX stall on BCM2712/RP1 (RFC patches from netdev)
* [`6a53a93`](https://github.com/siderolabs/pkgs/commit/6a53a933d1b9bf3c3e3d9fae8d7bc3e9021d418e) feat: bump kernel to 6.18.25
* [`f567bce`](https://github.com/siderolabs/pkgs/commit/f567bced2b6b5517cf70a5e925995e2dcdcd8444) feat: disable more stuff in Kconfig
* [`ffd9790`](https://github.com/siderolabs/pkgs/commit/ffd97909dd732c3ba8520ea4354ab2ecf07e8ba9) feat: bump kernel to 6.18.24
* [`b7c709a`](https://github.com/siderolabs/pkgs/commit/b7c709add255e09b3b1101abad06b4f3b17952cd) feat: bump deps
* [`e5e5b3c`](https://github.com/siderolabs/pkgs/commit/e5e5b3c0e65911069be6a62326fea677baac7245) feat: update Linux to 6.18.23
* [`1a4cd20`](https://github.com/siderolabs/pkgs/commit/1a4cd203fddcb04610bcf933c1f9058d94744863) fix: renovate config
* [`d0ed6ed`](https://github.com/siderolabs/pkgs/commit/d0ed6ed134c4aca27b4c8ef9dfc87476905487d4) feat: update dependencies
* [`6ea49c7`](https://github.com/siderolabs/pkgs/commit/6ea49c7264baf6948e8b793f0b8c1306f71efe5a) fix: support disabling module signature verification
* [`6520ec4`](https://github.com/siderolabs/pkgs/commit/6520ec481c215cbfcd44996e07cdb87057f12c71) feat: update containerd to 2.2.3
* [`37ce992`](https://github.com/siderolabs/pkgs/commit/37ce992e6a7d576fce9432fcf30fb7a656056d89) feat: enable CONFIG_UHID and CONFIG_INPUT_JOYDEV as modules
* [`cddd934`](https://github.com/siderolabs/pkgs/commit/cddd934ff6704bce64fe5861518d40801d6574f4) feat: update backportable dependencies
* [`32e4077`](https://github.com/siderolabs/pkgs/commit/32e4077a095576ac5b0f32fb08fd7601ccf4f30f) feat: update OpenSSL
* [`2d241e7`](https://github.com/siderolabs/pkgs/commit/2d241e7ec587a16fcf16aac8ad2ed47dfa38253b) feat: update Go to 1.26.2 and small deps updates
* [`7f540ce`](https://github.com/siderolabs/pkgs/commit/7f540ce7f367484cd44eb1d5ce25b59cf1cd1dce) feat: disable dynamic SCS
* [`3bef043`](https://github.com/siderolabs/pkgs/commit/3bef04361931a686d163a0c3cc76165f1059b838) feat: update runc to 1.4.2
* [`c6e6f10`](https://github.com/siderolabs/pkgs/commit/c6e6f1004e9f2947e0aea42a0baee197e745576f) feat: update Linux to 6.18.21
* [`a9e8afa`](https://github.com/siderolabs/pkgs/commit/a9e8afa610b325c5cbc6470bc62be92849dc5b88) fix: libarchive install prefix
* [`e4d0113`](https://github.com/siderolabs/pkgs/commit/e4d0113483e8c1920efc74037a10a82757493560) feat: update for musl 1.2.6
* [`9142603`](https://github.com/siderolabs/pkgs/commit/9142603113d8668de274b2cb207c69ae0a630e1c) feat: update NVIDIA production to 595.58.03
* [`22fa669`](https://github.com/siderolabs/pkgs/commit/22fa66967bf36b727a004495f0457049313be1f5) feat: update Linux to 6.18.19
* [`03680ae`](https://github.com/siderolabs/pkgs/commit/03680ae6e2e00501115415733a09891a5fd2fc35) feat: update containerd patch verifier role
* [`bdc239e`](https://github.com/siderolabs/pkgs/commit/bdc239e6a293bad5ba274874ceaf5f3d98a62284) feat: enable CHECKPOINT_RESTORE option
* [`83f5bcd`](https://github.com/siderolabs/pkgs/commit/83f5bcd9e3d02efc350356cbedb1c84acd00ca6e) chore: update toolchain and tools
* [`4f784de`](https://github.com/siderolabs/pkgs/commit/4f784de3801022a00691a4aac516b1f9a0b4d451) fix: install apparmor parser require config files
* [`559b1be`](https://github.com/siderolabs/pkgs/commit/559b1bee1299d621308d844128eede7875ed75e8) feat: enable AMD GPU peer-to-peer DMA
* [`77194e4`](https://github.com/siderolabs/pkgs/commit/77194e453b6e53392cbef215d0e78f2f75ebccb0) fix: disable CONFIG_RT_GROUP_SCHED
* [`02ee1e3`](https://github.com/siderolabs/pkgs/commit/02ee1e3a7bdc0a04ab44a40f938253c780eccd44) feat: backportable deps update
* [`21af1c3`](https://github.com/siderolabs/pkgs/commit/21af1c372f5a45be6ac18adac89935dc6311e07e) feat: bump deps
* [`6935f6f`](https://github.com/siderolabs/pkgs/commit/6935f6f2f5eca79706a9a3403178980cdc7faf7d) feat(kernel): enable CONFIG_USB_UHCI_HCD on amd64
* [`2c89e9f`](https://github.com/siderolabs/pkgs/commit/2c89e9fbfc5b09e161426b9584712e8c382b3bce) feat: update containerd to 2.2.2
* [`866939b`](https://github.com/siderolabs/pkgs/commit/866939b2e9d663446d0a86cba3d522146ae141ba) feat: update tools with LLVM 22.1
* [`13d00e0`](https://github.com/siderolabs/pkgs/commit/13d00e0b85949826e100f9397a750f762b606adf) feat: enable dynamic preemption support
* [`7d0cc32`](https://github.com/siderolabs/pkgs/commit/7d0cc32570c5a0b873e2fba70aa8d403bd5d5cd0) feat: update Linux 6.18.16, NVIDIA, ZFS
* [`ef3a7c8`](https://github.com/siderolabs/pkgs/commit/ef3a7c83b6f8c6106de6ebbc48eedecabf7187f4) feat: update Go to 1.26.1
* [`8148601`](https://github.com/siderolabs/pkgs/commit/8148601ca6a24f3bb39d0ce45e0e8754dc2c3bd4) feat: add containerd patch to verify images
* [`b7c7ab2`](https://github.com/siderolabs/pkgs/commit/b7c7ab22a9f9ce01f665086ef9d50eefa6d90cd1) feat: update Linux to 6.18.15
* [`830fbac`](https://github.com/siderolabs/pkgs/commit/830fbacedf58be78ce84a0347a9346aab8d68613) feat: enable CONFIG_USB_IPHETH kernel module
* [`adc1714`](https://github.com/siderolabs/pkgs/commit/adc1714e9b79cb6aff298384ed37473f2a81a77e) feat: update Linux to 6.18.14
* [`3c982f8`](https://github.com/siderolabs/pkgs/commit/3c982f8df278cf76a7fd421711eaf23bdbf3e948) chore: update deps
* [`d065c59`](https://github.com/siderolabs/pkgs/commit/d065c5993c5994bc855c2894b5d8ab671c98ee28) feat: update Linux firmware to 20260221
* [`773ea3a`](https://github.com/siderolabs/pkgs/commit/773ea3a035cf01d01228cb95993dec17df77dd2c) feat: update Linux to 6.18.13
* [`6ca02b3`](https://github.com/siderolabs/pkgs/commit/6ca02b3129118749b2da47c5cfd6f25c377c0360) fix: make udev rules read only
* [`520141c`](https://github.com/siderolabs/pkgs/commit/520141cd49b156e3db33496f187360acc85c3e1f) feat: enable kernel irq time accounting
* [`8f6df51`](https://github.com/siderolabs/pkgs/commit/8f6df518459513a107786b7020df3d9546d64e27) feat: enable CONFIG_HID_MULTITOUCH
* [`6934b50`](https://github.com/siderolabs/pkgs/commit/6934b5057f6996997420d43fbe620729c8cf22d5) feat: add patch for Cilium BPF verifier rejection by the kernel
* [`5760aa7`](https://github.com/siderolabs/pkgs/commit/5760aa774e043d121921304863d335db7e9e9adf) feat: enable MLX5 Scalable Functions and TC offload in kernel
* [`c0c8bc5`](https://github.com/siderolabs/pkgs/commit/c0c8bc56eb19aa4b4246c1813ab284b329ee9ffe) feat: enable CONFIG_DRM_ACCEL and IVPU on amd64
* [`b9cc39d`](https://github.com/siderolabs/pkgs/commit/b9cc39dcbbfb79141a644e614fb5e62da3fd93aa) feat: build kernel with Clang and ThinLTO, update Go to 1.26
* [`3327386`](https://github.com/siderolabs/pkgs/commit/33273866b175bd09a1d1bbfd47a41798537cb1b0) chore: drop mellanox-ofed
* [`9013985`](https://github.com/siderolabs/pkgs/commit/9013985d859828105d9e3cd60c08e44fc4e11d07) feat: update dependencies
* [`17196f5`](https://github.com/siderolabs/pkgs/commit/17196f595e7347e1f233cf5c9a1f16d90ce4e04d) feat: update NVIDIA LTS to 580.126.16
* [`8f53ad2`](https://github.com/siderolabs/pkgs/commit/8f53ad27bc7f70f4475cf71cabf7f86b1e20d794) feat: update Linux to 6.18.9
* [`eff5ba0`](https://github.com/siderolabs/pkgs/commit/eff5ba0d0e720ca4e1e2ed58c0719490b7f6826b) feat: enable ip6_gre
* [`605ac0d`](https://github.com/siderolabs/pkgs/commit/605ac0d9cbb88be263618a553bbb1be785af2e97) chore: update deps
* [`7670ff4`](https://github.com/siderolabs/pkgs/commit/7670ff45458bd39f5ca6076ba4eb65f0b68cf2e4) feat: enable NFT_BRIDGE config
* [`dc737a6`](https://github.com/siderolabs/pkgs/commit/dc737a68c470c9498ec11bde09196809355d2463) chore: update kernel
* [`9b118b3`](https://github.com/siderolabs/pkgs/commit/9b118b3d0fe7f0df06a069065b86ab307fef3375) chore: update deps
* [`a63c227`](https://github.com/siderolabs/pkgs/commit/a63c2276eea0013463487cebf95ee35a37c5d9f6) feat: update OpenSSL to v3.6.1
* [`da7ab57`](https://github.com/siderolabs/pkgs/commit/da7ab5776bd1a6c551bfc6fe5919114721da0e1f) feat: add px-fuse pkg
* [`553e0fb`](https://github.com/siderolabs/pkgs/commit/553e0fb70f076a8bc53e283253b30ff819e627ff) feat: enable dm-integrity
* [`15a3cdf`](https://github.com/siderolabs/pkgs/commit/15a3cdf54884d5169895a1ff46682373688ac5e2) feat: update Linux to 6.18.6
* [`b518a19`](https://github.com/siderolabs/pkgs/commit/b518a196de93dd33e70faaff2342f67acb7dc49b) feat: update dependencies
* [`1b4fbf5`](https://github.com/siderolabs/pkgs/commit/1b4fbf56b270d5669116fa0d8f91a3b9495e0d97) feat: update GRUB to 2.14
* [`30bc671`](https://github.com/siderolabs/pkgs/commit/30bc671d4be566ebf60b820edd54000616262e79) fix: enable pinctrl for Raspberry Pi 5
* [`375983f`](https://github.com/siderolabs/pkgs/commit/375983f4685484a8be5796f815629a9a0d8bd146) feat: update Go to 1.25.6
* [`d445c80`](https://github.com/siderolabs/pkgs/commit/d445c8076b7dd18b04f48e0a7e5cc2e50b3064d0) feat: update Linux to 6.18.5
* [`6994400`](https://github.com/siderolabs/pkgs/commit/69944002f9ee681220dcb23031c23ee327e6c1f2) feat: update NVIDIA LTS and production driver versions
* [`05c3d85`](https://github.com/siderolabs/pkgs/commit/05c3d856b7de6eb64af718d7266a5adf15e1224b) feat: update Linux firmware to 20260110
* [`c61b466`](https://github.com/siderolabs/pkgs/commit/c61b466e130015b44962e7ef3bc1e9bec935b1df) feat: enable IT87 hwmon module
* [`ae2572e`](https://github.com/siderolabs/pkgs/commit/ae2572e894a3d8d951418d447ec02f6cc65c8e72) feat: enable IPV6_MROUTE
* [`d6b503e`](https://github.com/siderolabs/pkgs/commit/d6b503e0fe75d52f83d656a3460cb3614b352e51) feat: add RK3588 NPU Support
* [`df4b4c8`](https://github.com/siderolabs/pkgs/commit/df4b4c885d4aabf702ce03bcb341f5b5f3641d76) feat: bump deps
* [`a220898`](https://github.com/siderolabs/pkgs/commit/a2208985bd756ef6366497c5f9768e814b3f7583) feat: add libarchive
* [`c2371b5`](https://github.com/siderolabs/pkgs/commit/c2371b5582836e27b3e80c4404c4ff5fbed90291) feat: enable ZRAM support
* [`ab4d169`](https://github.com/siderolabs/pkgs/commit/ab4d169ad93203ba56b0677a10e78eb3e623762e) feat: add a patch to force uid when populating from a directory
* [`972f44d`](https://github.com/siderolabs/pkgs/commit/972f44d5dae53809ef337544c52c835373439d34) feat: update dependencies
* [`f8eb5b0`](https://github.com/siderolabs/pkgs/commit/f8eb5b02aaebaf76c59e71f57f4a689dc727e769) feat: update Linux to 6.18.2
* [`3fb6291`](https://github.com/siderolabs/pkgs/commit/3fb629109a7e5f9650d0e641ff5076a29c319448) feat: update systemd to 259
* [`59241bd`](https://github.com/siderolabs/pkgs/commit/59241bd58eeb07a18af1c9fc8fffff6365ecca0d) fix: add SBOMs for pigz/igzip
* [`9377c78`](https://github.com/siderolabs/pkgs/commit/9377c786d112b4181f1e373f6e513130f11b7801) feat: optimize decompression for containerd
* [`e8e61ce`](https://github.com/siderolabs/pkgs/commit/e8e61cedbbd687ed958db992e05b5d59e4a8ea60) feat: update containerd to 2.2.1
* [`daa74ba`](https://github.com/siderolabs/pkgs/commit/daa74bab83f91bbc4b6c42625d2953299d5fe20a) feat: support xfs filesystem reproducibility
* [`1f66513`](https://github.com/siderolabs/pkgs/commit/1f665130fbda76478c261dd54e3843c15027c9cd) feat: update OpenZFS to 2.4.0
* [`b209af5`](https://github.com/siderolabs/pkgs/commit/b209af5baf1a67472ef431e5a8b7d48022392a1e) chore: rekres with latest changes
* [`2b806b9`](https://github.com/siderolabs/pkgs/commit/2b806b9b2a7e05b97c2a7e8572e3a8edbd3721d3) feat: bump dependencies
* [`65242fd`](https://github.com/siderolabs/pkgs/commit/65242fd0fef5c9c923aacce23d1655bad0d1b3e3) feat: enable CONFIG_MISC_RP1 in ARM64 config
* [`4daecd8`](https://github.com/siderolabs/pkgs/commit/4daecd8e7b8d87110a9e552a60a5394014294e08) feat: update Linux to 6.18.1
* [`9868a66`](https://github.com/siderolabs/pkgs/commit/9868a66e3c000f505c97ff68e61abac9c9e8e4c9) feat: enable Powercap and Intel RAPL
* [`07883ee`](https://github.com/siderolabs/pkgs/commit/07883eee3729d4d3adaaebcd825452934c3baebb) feat: build and package perf binary
* [`47abca0`](https://github.com/siderolabs/pkgs/commit/47abca0852b9555d88eba61661c65a7f93ec3590) fix: add json support to nftables binary
* [`b961ff8`](https://github.com/siderolabs/pkgs/commit/b961ff898fc9eae68d7f3cea2ca22ff4d0b9c99d) feat: patch containerd 2.2.0 with cgroups fix patch
* [`b7dd7f6`](https://github.com/siderolabs/pkgs/commit/b7dd7f6c809f670f058b78fd3b84f4cb977771cb) feat: add mstflint module
* [`ae53351`](https://github.com/siderolabs/pkgs/commit/ae5335198e009da7b06bc0f0d6f42b0947650fc0) feat: update ZFS to 2.4.0-rc5
* [`b8edf01`](https://github.com/siderolabs/pkgs/commit/b8edf0168171ffc5b87fcd962e37d5c2cd25b687) feat: update CNI plugins to v1.9.0
* [`a57c1b0`](https://github.com/siderolabs/pkgs/commit/a57c1b0c9d143559a87b64fe9570eec39c14a771) feat: enable amd sev-snp
* [`68562c1`](https://github.com/siderolabs/pkgs/commit/68562c1b4cdba656287021a1694440b2a7e4d24d) feat: update Linux to 6.18
* [`6f4ff8c`](https://github.com/siderolabs/pkgs/commit/6f4ff8cc9f57452707588c05e5ca4e80c56548d2) feat: enable Amlogic Meson PCIe controller driver
* [`c41127b`](https://github.com/siderolabs/pkgs/commit/c41127b94d22b9a5cb6b93f49b546f2ff477410c) feat: enable Intel GPIO/Pinctrl kernel modules
* [`4a31ff7`](https://github.com/siderolabs/pkgs/commit/4a31ff7dd5c9266b68abded53a7399cb8102f4e3) feat: update NVIDIA LTS to 580.105.08
</p>
</details>

### Changes from siderolabs/tools
<details><summary>50 commits</summary>
<p>

* [`80fd6eb`](https://github.com/siderolabs/tools/commit/80fd6ebfcac67f696003340e1e15aafd49b3378b) feat: add python pyyaml tools
* [`c59cce5`](https://github.com/siderolabs/tools/commit/c59cce5a50aa99d0f553f8d16dada00b64d8fb9f) feat: enable static libs for pcre2
* [`da6e92a`](https://github.com/siderolabs/tools/commit/da6e92aa5937a7b3b94b72aef54c788d8c8931fd) feat: use kernel gnu mirrors
* [`08071b1`](https://github.com/siderolabs/tools/commit/08071b1368ea5b1033212fe65b220ad4233d95d4) feat: update dependencies
* [`2723c06`](https://github.com/siderolabs/tools/commit/2723c06a8d0d4cb671eb1c6de64e4ae35231f865) chore: bump pkgfile bldr to v0.6.1
* [`d333b32`](https://github.com/siderolabs/tools/commit/d333b325104d5daf6f4f3665d95fb85b32048d63) chore: bump toolchain to latest
* [`c7494c6`](https://github.com/siderolabs/tools/commit/c7494c66a5f608ba79d836dd552404786c352af7) chore: deps 2026-07-13 (major bumps only)
* [`57a0183`](https://github.com/siderolabs/tools/commit/57a01830e10d93abde3bf230bdd78fb7fbc9e88c) chore: deps 2026-07-13 (minor bumps only)
* [`21db13e`](https://github.com/siderolabs/tools/commit/21db13ecfcf165c82f33bd2ac9aaf0018155010d) chore: deps 2026-07-13 (patch only)
* [`7079a8a`](https://github.com/siderolabs/tools/commit/7079a8a03e59765c8414367826d0ffc47838f340) chore: bump toolchain
* [`5326524`](https://github.com/siderolabs/tools/commit/5326524e363691bf6145d66ddb4cb0ce9462a715) feat: build LLVM as cross-compiling
* [`878f1db`](https://github.com/siderolabs/tools/commit/878f1db715a91322d0b199a539ab183d11e8d0cc) feat: bump dependencies
* [`0f1c859`](https://github.com/siderolabs/tools/commit/0f1c859d7e22e8534ad5de6c6ea99ffbb77a1400) chore: make rekres
* [`5c0c9be`](https://github.com/siderolabs/tools/commit/5c0c9be532cadebc2a94346f1c1a1437c3637be3) chore: update dependencies 2026-06-16
* [`b88d99c`](https://github.com/siderolabs/tools/commit/b88d99c4d556f8b715502bdac873e627a863e9ec) feat: bump OpenSSL to 3.6.3
* [`42c59b9`](https://github.com/siderolabs/tools/commit/42c59b990ed6e95fa3dda7347b7fa5684c37fbb3) feat: bump toolchain to bring in Go 1.26.4
* [`206a4c0`](https://github.com/siderolabs/tools/commit/206a4c0ea091efa8939247069cfeee712da7c077) feat: update dependencies, rework LLVM build
* [`f9f37df`](https://github.com/siderolabs/tools/commit/f9f37df3c44979542901d88690ce93108ef6a037) fix: add proper name for zlib-ng sbom
* [`aa45c41`](https://github.com/siderolabs/tools/commit/aa45c41b733c450b21d7503a978c8b8044ebc2fa) fix: add SBOM for zstd library
* [`808f34f`](https://github.com/siderolabs/tools/commit/808f34f620e698863c509595a60ecf192310bfb1) feat: update Go to 1.26.3
* [`5dfe83d`](https://github.com/siderolabs/tools/commit/5dfe83d244c4c4a6cc104b4f45cd4f76ae2873d8) feat: drop fakeroot and policycoreutils
* [`618fd20`](https://github.com/siderolabs/tools/commit/618fd2061fe885a54014cbe1222cbd3c95471807) feat: add Python wheel package
* [`df3c1b7`](https://github.com/siderolabs/tools/commit/df3c1b7bde417e33e67f4fb43494ff1ecd0b3399) feat: bump dependencies
* [`44ad18c`](https://github.com/siderolabs/tools/commit/44ad18c5a553eb2f728f369a8c56e3c257730da2) feat: bump deps
* [`f3d0dd9`](https://github.com/siderolabs/tools/commit/f3d0dd9ca5c9006ca14890af1ab8a58248ae28d8) fix: renovate configs
* [`4ac4449`](https://github.com/siderolabs/tools/commit/4ac444995923055b5c410dc957579f4b0b308394) feat: update dependencies
* [`027744f`](https://github.com/siderolabs/tools/commit/027744f476f38f0fda9b1fd0ae7fb3aed0ab4ad1) feat: bump OpenSSL to 3.6.2
* [`7067f1f`](https://github.com/siderolabs/tools/commit/7067f1f966cff98c83cf2a4ecfaf06021397d954) feat: update util-linux to 2.41.4
* [`6cb3e56`](https://github.com/siderolabs/tools/commit/6cb3e561ff60abc78cefd570189651c8afdc7121) feat: update Go to 1.26.2
* [`9186c5f`](https://github.com/siderolabs/tools/commit/9186c5ffff2bffa4b92d7377d254faedceba6036) feat: update musl to 1.2.6
* [`6d7a2e4`](https://github.com/siderolabs/tools/commit/6d7a2e42f0608dec3a791c0b37cd6c0781295776) chore: update toolchain
* [`5c753fc`](https://github.com/siderolabs/tools/commit/5c753fc00c75d742638ab870b31acc86eab4e9ab) feat: bump deps
* [`4075c05`](https://github.com/siderolabs/tools/commit/4075c05e401f2abb049bdd97ff7a16dadd4eebff) feat: update LLVM to 22
* [`f175a91`](https://github.com/siderolabs/tools/commit/f175a918f8371b02cf486dba533bb981439aec57) feat: update Go to 1.26.1 and other bumps
* [`9de9770`](https://github.com/siderolabs/tools/commit/9de9770aeffac63ebe0cbd9e02ebe8f77b4a9635) feat: update to Go 1.26
* [`bd4ae8f`](https://github.com/siderolabs/tools/commit/bd4ae8f69c218c05e143a1a04c5710505b8559f2) feat: add LLVM+Clang+LLD toolchain
* [`90bd70c`](https://github.com/siderolabs/tools/commit/90bd70c94cf434ff45dcd71d1f3d9435b8243093) feat: update dependencies
* [`decb988`](https://github.com/siderolabs/tools/commit/decb9887929e8d3b3879d56a984244cfe8ae0213) chore: update deps
* [`ca26e1c`](https://github.com/siderolabs/tools/commit/ca26e1c38cd0a76eb981db4dad2e6caccb0bbe4d) chore: update deps
* [`0281af0`](https://github.com/siderolabs/tools/commit/0281af0545e17c409fb32d8db61a7d9b0ad8b1c2) feat: update OpenSSL to 3.6.1
* [`721ad07`](https://github.com/siderolabs/tools/commit/721ad073f18b41407882727a3f0061e594f6c955) feat: update dependencies
* [`2b3f514`](https://github.com/siderolabs/tools/commit/2b3f514d42a343d98c79a487e80bd4f225a41b70) fix: reproducible build for nasm
* [`98c699e`](https://github.com/siderolabs/tools/commit/98c699eb624d0846455f08db77cc14e446cb6db9) feat: update Go to 1.25.6
* [`cd5eb66`](https://github.com/siderolabs/tools/commit/cd5eb66bb0de4fb468a860e176267c3420b4a3a1) chore: run rekres and update dependencies
* [`896f8b9`](https://github.com/siderolabs/tools/commit/896f8b9c1f88cd190d11b8ef3baa2c36e73d6dfe) fix: add sbom for zlib-ng
* [`543a16f`](https://github.com/siderolabs/tools/commit/543a16fedf7170d8b015ea1391817328205e629a) feat: replace zlib -> zlib-ng, add nasm
* [`b67c1a1`](https://github.com/siderolabs/tools/commit/b67c1a168b302539d2082a5513c4a0130c30e4df) chore: rekres with latest changes
* [`5e087cb`](https://github.com/siderolabs/tools/commit/5e087cbcd158db1ce4f447145bd76a24d07159a1) feat: bump dependencies
* [`da96a27`](https://github.com/siderolabs/tools/commit/da96a2771801627b4715f7a13199aa6846f87732) chore: rekres to fix reproducibility
* [`e283ec8`](https://github.com/siderolabs/tools/commit/e283ec8d3831bb19b26938afb10f4955ea563ce2) feat: update Go to 1.25.5
</p>
</details>

### Dependency Changes

* **github.com/siderolabs/pkgs**   v1.13.0-alpha.0 -> v1.14.0-alpha.0-127-g7d8b87b
* **github.com/siderolabs/tools**  v1.13.0-alpha.0 -> v1.14.0-alpha.0-29-g80fd6eb

Previous release can be found at [v0.2.0](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.2.0)

## [sbc-raspberrypi 0.2.0](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.2.0) (2026-02-25)

Welcome to the v0.2.0 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Noel Georgi

### Changes
<details><summary>1 commit</summary>
<p>

* [`0a02d01`](https://github.com/siderolabs/sbc-raspberrypi/commit/0a02d01d115c9daf18134ae34699d7364637963a) feat: bump tools and pkgs
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.9](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.9)

## [sbc-raspberrypi 0.1.9](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.9) (2026-02-24)

Welcome to the v0.1.9 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Birger Johan Nordølum
* Licia Seiker
* Mateusz Urbanek

### Changes
<details><summary>3 commits</summary>
<p>

* [`3f8fd68`](https://github.com/siderolabs/sbc-raspberrypi/commit/3f8fd6866a0295e26315a08931555f19c480c42c) chore: update kres's renovate package rules
* [`e56ea57`](https://github.com/siderolabs/sbc-raspberrypi/commit/e56ea57af329f5a287ba17420a132af5f5f21c96) feat: add DTB build for Raspberry Pi 5
* [`bb1434a`](https://github.com/siderolabs/sbc-raspberrypi/commit/bb1434abb62a413dcf73ba39b4bffad93b02512d) chore: update deps
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.8](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.8)

## [sbc-raspberrypi 0.1.8](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.8) (2026-01-26)

Welcome to the v0.1.8 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Raspberry Pi 5 Support



### Contributors

* Dmitrii Sharshakov

### Changes
<details><summary>2 commits</summary>
<p>

* [`f35948c`](https://github.com/siderolabs/sbc-raspberrypi/commit/f35948c0d7fcf0adbd381d55bfae4d91379ae488) feat: add Raspberry Pi 5
* [`9a3783b`](https://github.com/siderolabs/sbc-raspberrypi/commit/9a3783b361d7d3f79d9ac9bc272f9dd000c3f729) fix: replace Pi 5 dts with ones from the kernel package
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.7](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.7)

## [sbc-raspberrypi 0.1.7](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.7) (2025-12-16)

Welcome to the v0.1.7 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Mateusz Urbanek
* Oscar Virot
* eseiker

### Changes
<details><summary>3 commits</summary>
<p>

* [`1e01fd5`](https://github.com/siderolabs/sbc-raspberrypi/commit/1e01fd5f6d542faebfc82a6c03d5493438875f88) fix: ensure we apply [all] filters to config
* [`09feadd`](https://github.com/siderolabs/sbc-raspberrypi/commit/09feadd74a78dd4cba5d96018032bb55df4b9d8c) chore: update deps
* [`b2390bd`](https://github.com/siderolabs/sbc-raspberrypi/commit/b2390bdb453802613ea1fa44bc60437851b1f133) fix: disable UART for Raspberry Pi 5 compatibility
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.6](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.6)

## [sbc-raspberrypi 0.1.6](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.6) (2025-11-13)

Welcome to the v0.1.6 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Noel Georgi
* eseiker

### Changes
<details><summary>3 commits</summary>
<p>

* [`a41e85b`](https://github.com/siderolabs/sbc-raspberrypi/commit/a41e85ba7eb09747b569f9621071a9905665bd52) chore: bump u-boot to support raspberry pi 5
* [`bb2ccbe`](https://github.com/siderolabs/sbc-raspberrypi/commit/bb2ccbe805fa40583e4f5939ef069ae06d03a825) chore: bump tools
* [`de6eaa9`](https://github.com/siderolabs/sbc-raspberrypi/commit/de6eaa9c4d5b14a5c53435615984071ceaed92ae) chore: bump deps
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.5](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.5)

## [sbc-raspberrypi 0.1.5](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.5) (2025-05-02)

Welcome to the v0.1.5 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Noel Georgi

### Changes
<details><summary>1 commit</summary>
<p>

* [`666e16e`](https://github.com/siderolabs/sbc-raspberrypi/commit/666e16ee039789614442dc2bcf27001d20e3ab34) fix: use grub as bootloader for SBCs
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.4](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.4)

## [sbc-raspberrypi 0.1.4](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.4) (2025-04-15)

Welcome to the v0.1.4 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Noel Georgi
* Andrey Smirnov
* Martin Schuessler

### Changes
<details><summary>4 commits</summary>
<p>

* [`8447811`](https://github.com/siderolabs/sbc-raspberrypi/commit/8447811443f76fb9da1e993842ccadfade886bfe) chore: bump deps
* [`71bb718`](https://github.com/siderolabs/sbc-raspberrypi/commit/71bb718d7c257dbf2dd3e11a3e4ba33c17b3bd7f) feat: add support for the revolutionpi
* [`3eeedb3`](https://github.com/siderolabs/sbc-raspberrypi/commit/3eeedb3ee5d8bd12e56129c30c26cc1b2884a379) feat: use new tools as base
* [`2bb6b52`](https://github.com/siderolabs/sbc-raspberrypi/commit/2bb6b5299af42e2b48ee3d7964cabfb7f46c8fd6) chore: unify buildkits
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.3](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.3)

## [sbc-raspberrypi 0.1.3](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.3) (2025-02-04)

Welcome to the v0.1.3 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Andrey Smirnov

### Changes
<details><summary>1 commit</summary>
<p>

* [`cbb744e`](https://github.com/siderolabs/sbc-raspberrypi/commit/cbb744ed5da92a22bb6578ebeadfdd8e41837554) feat: revert enable watchdog device in firmware by default
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.2](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.2)

## [sbc-raspberrypi 0.1.2](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.2) (2025-01-28)

Welcome to the v0.1.2 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Nico Berlee

### Changes
<details><summary>1 commit</summary>
<p>

* [`e113e8b`](https://github.com/siderolabs/sbc-raspberrypi/commit/e113e8b12749ca309903022fa6eafe60d14a7919) feat: enable watchdog device in firmware by default
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.1](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.1)

## [sbc-raspberrypi 0.1.1](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.1) (2024-12-23)

Welcome to the v0.1.1 release of sbc-raspberrypi!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Nico Berlee
* Noel Georgi

### Changes
<details><summary>2 commits</summary>
<p>

* [`a802dbb`](https://github.com/siderolabs/sbc-raspberrypi/commit/a802dbb773b927e673b76023faf386bf9f74c281) fix: fix raspberrypi firmware url
* [`2f30031`](https://github.com/siderolabs/sbc-raspberrypi/commit/2f300313aa7f277b194a21e1326d29b97f17372e) chore: rekres to simplify `.kres.yaml` defaults
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.0](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0)

## [sbc-raspberrypi 0.1.0](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0) (2024-12-02)

Welcome to the v0.1.0 release of sbc-raspberrypi!  
*This is a pre-release of sbc-raspberrypi*



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Noel Georgi
* alphabet5

### Changes
<details><summary>2 commits</summary>
<p>

* [`15a6834`](https://github.com/siderolabs/sbc-raspberrypi/commit/15a683428278f8bf1b7c46110e6cb72d417ad1b8) chore: bump deps
* [`1c69056`](https://github.com/siderolabs/sbc-raspberrypi/commit/1c690560d5a16cefc5977f033dae102424f2ea63) docs: add link to boot assets
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.0-beta.2](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0-beta.2)

## [sbc-raspberrypi 0.1.0-beta.2](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0-beta.2) (2024-08-08)

Welcome to the v0.1.0-beta.2 release of sbc-raspberrypi!  
*This is a pre-release of sbc-raspberrypi*



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Noel Georgi

### Changes
<details><summary>1 commit</summary>
<p>

* [`cb7b03d`](https://github.com/siderolabs/sbc-raspberrypi/commit/cb7b03d20a390486f4e46be43a4d6aacefac734d) chore: rekres and bump deps
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.0-beta.1](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0-beta.1)

## [sbc-raspberrypi 0.1.0-beta.1](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0-beta.1) (2024-07-05)

Welcome to the v0.1.0-beta.1 release of sbc-raspberrypi!  
*This is a pre-release of sbc-raspberrypi*



Please try out the release binaries and report any issues at
https://github.com/siderolabs/sbc-raspberrypi/issues.

### Contributors

* Andrey Smirnov
* Matthias Riegler

### Changes
<details><summary>3 commits</summary>
<p>

* [`3967840`](https://github.com/siderolabs/sbc-raspberrypi/commit/396784073bbf066bb1243ef2934e86913f147f4c) release(v0.1.0-beta.1): prepare release
* [`2590781`](https://github.com/siderolabs/sbc-raspberrypi/commit/2590781a2afb4f333f43732aa9c42b717bd173ee) feat: update dependencies
* [`9bb942b`](https://github.com/siderolabs/sbc-raspberrypi/commit/9bb942b21f6be9bf104386cf6502ab50af5e2ae8) fix: nvme boot is only considered when network disconnected
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v0.1.0-beta.0](https://github.com/siderolabs/sbc-raspberrypi/releases/tag/v0.1.0-beta.0)

