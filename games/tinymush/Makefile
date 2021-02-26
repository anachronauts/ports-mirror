# $FreeBSD$

PORTNAME=	tinymush
DISTVERSIONPREFIX=	v
DISTVERSION=	3.3.0.12-26
DISTVERSIONSUFFIX=	-g146b537
CATEGORIES=	games

USE_GITHUB=	yes
GH_ACCOUNT=	TinyMUSH
GH_PROJECT=	TinyMUSH

MAINTAINER=	sage@anachronauts.club
COMMENT=	A MUSH server

LIB_DEPENDS=	libltdl.so:devel/libltdl \
		libpcre.so.1:devel/pcre \
		libgdbm.so:databases/gdbm \
		libintl.so:devel/gettext-runtime

BUILD_DEPENDS=	cproto:devel/cproto

GNU_CONFIGURE=	yes
CONFIGURE_ARGS=	--localstatedir=/var --with-install-suffix=tinymush \
		--without-included-ltdl \
		--with-ltdl-lib=${PREFIX}/lib \
		--with-ltdl-include=${PREFIX}/include

USE_RC_SUBR=	tinymush

post-install:
.for config in alias compat netmush
	(cd ${STAGEDIR}${PREFIX}/etc/tinymush && cp ${config}.conf ${config}.conf.sample)
.endfor

.include <bsd.port.mk>
