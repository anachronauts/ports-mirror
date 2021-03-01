# $FreeBSD$

PORTNAME=	gmni
DISTVERSION=	g20210223
PORTREVISION=	0
CATEGORIES=	www

MAINTAINER=	sage@anachronauts.club
COMMENT=	Simple Gemini client

MASTER_SITES=	https://git.sr.ht/~sircmpwn/gmni/archive/
DISTFILES=	863c41db.tar.gz
WRKSRC=		${WRKDIR}/${PORTNAME}-863c41db

USES=		cmake
USE_LDCONFIG=	yes

post-patch:
	cp files/CMakeLists.txt ${WRKSRC}
	cp files/libgmni.pc.in ${WRKSRC}

.include <bsd.port.mk>
