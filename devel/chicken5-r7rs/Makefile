PORTNAME=	r7rs
DISTVERSION=	1.0.5
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Support for most of the R7RS Scheme language

LICENSE=        BSD3CLAUSE

MY_DEPENDS=	${PREFIX}/lib/chicken5/11/matchable.egg-info:devel/chicken5-matchable \
		${PREFIX}/lib/chicken5/11/srfi-1.egg-info:devel/chicken5-srfi-1 \
		${PREFIX}/lib/chicken5/11/srfi-13.egg-info:devel/chicken5-srfi-13
RUN_DEPENDS=	chicken5:lang/chicken5 \
		${MY_DEPENDS}
BUILD_DEPENDS=	chicken-install5:lang/chicken5 \
		${MY_DEPENDS}

CSC_OPTIONS=	-v -strip

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && DESTDIR=${STAGEDIR} chicken-install5 \
		-v -no-install-dependencies

.export CSC_OPTIONS
.include <bsd.port.mk>
