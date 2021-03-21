# $FreeBSD$

PORTNAME=	gmni
DISTVERSION=	g20210316
PORTREVISION=	0
CATEGORIES=	www

MAINTAINER=	sage@anachronauts.club
COMMENT=	Simple Gemini client

MASTER_SITES=	https://git.sr.ht/~sircmpwn/gmni/archive/
DISTFILES=	77b73efb.tar.gz
WRKSRC=		${WRKDIR}/${PORTNAME}-77b73efb

USES=		cmake
LIB_DEPENDS=	libbearssl.so:security/bearssl
USE_LDCONFIG=	yes

post-patch:
	cp files/CMakeLists.txt ${WRKSRC}
	cp files/libgmni.pc.in ${WRKSRC}

.include <bsd.port.mk>
