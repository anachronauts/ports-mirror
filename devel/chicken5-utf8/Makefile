# $FreeBSD$

PORTNAME=	utf8
DISTVERSION=	3.6.2
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Unicode support for CHICKEN

MY_DEPENDS=	${PREFIX}/lib/chicken5/11/iset.egg-info:devel/chicken5-iset \
		${PREFIX}/lib/chicken5/11/srfi-69.egg-info:devel/chicken5-srfi-69 \
		${PREFIX}/lib/chicken5/11/regex.egg-info:devel/chicken5-regex
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
