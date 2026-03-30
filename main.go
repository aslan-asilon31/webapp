package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="id">
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>ASLAN — Premium Store</title>
<link href="https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,400;0,700;0,900;1,400&family=Outfit:wght@300;400;500;600;700&display=swap" rel="stylesheet"/>
<style>
:root{
  --ink:#07070f;--deep:#0d0d1c;--card:#12121e;--panel:#181828;
  --border:rgba(255,255,255,.07);--border2:rgba(255,255,255,.12);
  --gold:#c8a45a;--gold2:#e8c87a;--goldf:linear-gradient(135deg,#c8a45a,#f0d080,#c8a45a);
  --teal:#00c8b0;--teal2:#00ffdd;
  --purple:#7c5cf5;--purplef:linear-gradient(135deg,#7c5cf5,#a78bfa);
  --text:#eeeef8;--sub:#8888a8;--muted:#55556a;
  --red:#ff4d70;--green:#22d3a0;
  --r:14px;--r2:22px;--nav:68px;
}
*,*::before,*::after{box-sizing:border-box;margin:0;padding:0}
html{scroll-behavior:smooth}
body{font-family:'Outfit',sans-serif;background:var(--ink);color:var(--text);min-height:100vh;overflow-x:hidden}
img{display:block;max-width:100%}a{color:inherit;text-decoration:none}button{cursor:pointer;font-family:'Outfit',sans-serif}
#amb{position:fixed;inset:0;z-index:0;pointer-events:none;overflow:hidden}
#amb span{position:absolute;border-radius:50%;filter:blur(120px);opacity:.15;animation:drift 18s ease-in-out infinite alternate}
#amb span:nth-child(1){width:600px;height:600px;background:#7c5cf5;top:-200px;left:-150px;animation-delay:0s}
#amb span:nth-child(2){width:450px;height:450px;background:#00c8b0;bottom:-100px;right:-100px;animation-delay:-7s}
#amb span:nth-child(3){width:350px;height:350px;background:#c8a45a;top:40%;left:55%;animation-delay:-13s}
@keyframes drift{0%{transform:translate(0,0) scale(1)}100%{transform:translate(50px,30px) scale(1.12)}}
nav{position:fixed;top:0;left:0;right:0;z-index:500;height:var(--nav);display:flex;align-items:center;justify-content:space-between;padding:0 36px;background:rgba(7,7,15,.75);backdrop-filter:blur(24px);border-bottom:1px solid var(--border)}
.nav-logo{font-family:'Playfair Display',serif;font-size:1.5rem;font-weight:900;letter-spacing:3px;background:var(--goldf);-webkit-background-clip:text;-webkit-text-fill-color:transparent}
.nav-links{display:flex;gap:4px}
.nav-links a{padding:8px 16px;border-radius:10px;font-size:.875rem;font-weight:500;color:var(--sub);transition:all .2s;position:relative}
.nav-links a:hover,.nav-links a.active{color:var(--text);background:rgba(255,255,255,.06)}
.nav-links a.active::after{content:'';position:absolute;bottom:5px;left:50%;transform:translateX(-50%);width:14px;height:2px;background:var(--gold);border-radius:2px}
.nav-right{display:flex;align-items:center;gap:10px}
.cart-btn{position:relative;width:40px;height:40px;border-radius:10px;border:1px solid var(--border2);background:rgba(255,255,255,.05);display:flex;align-items:center;justify-content:center;font-size:1.1rem;transition:all .2s}
.cart-btn:hover{background:rgba(255,255,255,.1);border-color:var(--gold)}
.cart-ct{position:absolute;top:-7px;right:-7px;background:var(--gold);color:#000;font-size:.62rem;font-weight:700;width:19px;height:19px;border-radius:50%;display:none;align-items:center;justify-content:center}
.ham{display:none;flex-direction:column;gap:5px;padding:8px;cursor:pointer}
.ham span{width:22px;height:2px;background:var(--text);border-radius:2px;transition:all .3s}
.mmenu{display:none;position:fixed;inset:0;z-index:490;background:var(--deep);padding:calc(var(--nav) + 24px) 24px 24px;flex-direction:column;gap:8px}
.mmenu.open{display:flex}
.mmenu a{padding:14px 20px;border-radius:12px;font-size:1rem;font-weight:500;color:var(--sub);transition:all .2s;border:1px solid transparent}
.mmenu a:hover,.mmenu a.active{color:var(--text);background:var(--card);border-color:var(--border2)}
.page{display:none;min-height:100vh;padding-top:var(--nav)}
.page.active{display:block}
.btn{padding:12px 28px;border:none;border-radius:12px;font-size:.9rem;font-weight:600;letter-spacing:.3px;transition:all .22s}
.btn-gold{background:var(--goldf);color:#0a0a14;box-shadow:0 4px 20px rgba(200,164,90,.3)}
.btn-gold:hover{transform:translateY(-2px);box-shadow:0 8px 28px rgba(200,164,90,.45)}
.btn-outline{background:transparent;border:1.5px solid var(--border2);color:var(--text)}
.btn-outline:hover{border-color:var(--gold);color:var(--gold)}
.btn-ghost{background:rgba(255,255,255,.06);border:1px solid var(--border);color:var(--text)}
.btn-ghost:hover{background:rgba(255,255,255,.1)}
.btn-teal{background:linear-gradient(135deg,var(--teal),var(--teal2));color:#0a0a14}
.btn-teal:hover{transform:translateY(-2px);box-shadow:0 6px 20px rgba(0,200,176,.35)}
.btn-danger{background:linear-gradient(135deg,#e05,#ff4d70);color:#fff}
.btn-sm{padding:8px 18px;font-size:.8rem}
.stag{display:inline-block;font-size:.72rem;font-weight:700;letter-spacing:2px;text-transform:uppercase;color:var(--gold);margin-bottom:10px}
.stitle{font-family:'Playfair Display',serif;font-size:clamp(1.8rem,4vw,3rem);font-weight:700;line-height:1.2}
.ssub{color:var(--sub);font-size:.98rem;margin-top:10px;max-width:540px;line-height:1.7}
/* HERO */
.hero{position:relative;min-height:calc(100vh - var(--nav));display:flex;align-items:center;padding:80px 36px 60px;overflow:hidden}
.hero-content{position:relative;z-index:2;max-width:580px}
.hero-eyebrow{display:flex;align-items:center;gap:10px;margin-bottom:22px}
.hero-eyebrow span{font-size:.73rem;font-weight:600;letter-spacing:2px;text-transform:uppercase;color:var(--gold);padding:5px 14px;border-radius:99px;border:1px solid rgba(200,164,90,.3);background:rgba(200,164,90,.08)}
.hero h1{font-family:'Playfair Display',serif;font-size:clamp(2.6rem,6vw,5rem);font-weight:900;line-height:1.06;letter-spacing:-1px}
.hero h1 em{font-style:italic;background:var(--goldf);-webkit-background-clip:text;-webkit-text-fill-color:transparent}
.hero-desc{color:var(--sub);font-size:1rem;line-height:1.78;margin:22px 0 34px;max-width:460px}
.hero-actions{display:flex;gap:14px;flex-wrap:wrap}
.hero-stats{display:flex;gap:36px;margin-top:56px;padding-top:36px;border-top:1px solid var(--border);flex-wrap:wrap}
.hstat .num{font-family:'Playfair Display',serif;font-size:1.9rem;font-weight:700;background:var(--goldf);-webkit-background-clip:text;-webkit-text-fill-color:transparent}
.hstat .lbl{font-size:.75rem;color:var(--sub);margin-top:2px;letter-spacing:.5px}
.hero-vis{position:absolute;right:-40px;top:50%;transform:translateY(-50%);width:min(580px,50vw);z-index:1}
.hero-img{width:100%;border-radius:24px;object-fit:cover;box-shadow:0 40px 100px rgba(0,0,0,.65);aspect-ratio:4/5}
.hfc{position:absolute;background:rgba(18,18,30,.92);backdrop-filter:blur(20px);border:1px solid var(--border2);border-radius:16px;padding:13px 18px;animation:float 4s ease-in-out infinite}
.hfc:nth-child(2){bottom:50px;left:-70px;animation-delay:0s}
.hfc:nth-child(3){top:40px;right:-30px;animation-delay:-2s}
@keyframes float{0%,100%{transform:translateY(0)}50%{transform:translateY(-10px)}}
.fl-lbl{font-size:.66rem;color:var(--sub);letter-spacing:.5px;margin-bottom:3px}
.fl-val{font-weight:700;font-size:.93rem;color:var(--gold)}
/* MARQUEE */
.mq{border-top:1px solid var(--border);border-bottom:1px solid var(--border);padding:14px 0;overflow:hidden;background:rgba(255,255,255,.015)}
.mq-track{display:flex;gap:48px;white-space:nowrap;animation:marq 22s linear infinite}
.mq-track:hover{animation-play-state:paused}
.mq-item{font-size:.75rem;font-weight:600;letter-spacing:2px;text-transform:uppercase;color:var(--muted);flex-shrink:0}
.mq-item span{color:var(--gold);margin-right:48px}
@keyframes marq{0%{transform:translateX(0)}100%{transform:translateX(-50%)}}
/* FEATURED */
.featured{padding:90px 36px;max-width:1280px;margin:0 auto}
.pgrid{display:grid;grid-template-columns:repeat(auto-fill,minmax(250px,1fr));gap:22px;margin-top:46px}
.pcard{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);overflow:hidden;transition:all .3s;cursor:pointer;position:relative}
.pcard:hover{transform:translateY(-6px);border-color:rgba(200,164,90,.3);box-shadow:0 20px 60px rgba(0,0,0,.45)}
.pbadge{position:absolute;top:13px;left:13px;z-index:2;font-size:.66rem;font-weight:700;letter-spacing:.8px;text-transform:uppercase;padding:3px 10px;border-radius:6px}
.b-new{background:rgba(0,200,176,.15);color:var(--teal);border:1px solid rgba(0,200,176,.3)}
.b-sale{background:rgba(255,77,112,.15);color:var(--red);border:1px solid rgba(255,77,112,.3)}
.b-hot{background:rgba(200,164,90,.15);color:var(--gold);border:1px solid rgba(200,164,90,.3)}
.pimg{width:100%;aspect-ratio:1;object-fit:cover;background:var(--panel)}
.pbody{padding:16px}
.pcat{font-size:.7rem;color:var(--sub);letter-spacing:.8px;text-transform:uppercase;margin-bottom:5px}
.pname{font-weight:600;font-size:.93rem;line-height:1.4;margin-bottom:10px;display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden}
.pbottom{display:flex;align-items:center;justify-content:space-between}
.pprice{font-family:'Playfair Display',serif;font-size:1.05rem;font-weight:700;color:var(--gold)}
.prating{font-size:.75rem;color:var(--sub)}
.padd{width:34px;height:34px;border-radius:9px;border:1px solid var(--border2);background:transparent;display:flex;align-items:center;justify-content:center;font-size:1.1rem;transition:all .2s;color:var(--text)}
.padd:hover{background:var(--gold);border-color:var(--gold);color:#000}
/* CATEGORIES */
.categories{padding:0 36px 90px;max-width:1280px;margin:0 auto}
.cgrid{display:grid;grid-template-columns:repeat(auto-fill,minmax(150px,1fr));gap:14px;margin-top:38px}
.citem{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:26px 18px;text-align:center;cursor:pointer;transition:all .25s}
.citem:hover{border-color:rgba(200,164,90,.4);background:var(--panel);transform:translateY(-3px)}
.cicon{font-size:1.9rem;margin-bottom:10px}
.cname{font-size:.83rem;font-weight:600;color:var(--sub)}
.citem:hover .cname{color:var(--gold)}
/* PROMO */
.promo{padding:0 36px 90px;max-width:1280px;margin:0 auto}
.promo-in{background:linear-gradient(135deg,#1a1040 0%,#0d1a30 50%,#0a1a20 100%);border-radius:var(--r2);padding:56px;display:flex;align-items:center;justify-content:space-between;gap:36px;border:1px solid rgba(124,92,245,.2);position:relative;overflow:hidden;flex-wrap:wrap}
.promo-in::before{content:'';position:absolute;inset:0;background:radial-gradient(ellipse 60% 80% at 70% 50%,rgba(0,200,176,.07),transparent)}
.promo-txt{position:relative;z-index:1;max-width:500px}
.promo-txt h2{font-family:'Playfair Display',serif;font-size:clamp(1.5rem,3vw,2.3rem);font-weight:700;margin-bottom:10px}
.promo-txt p{color:var(--sub);line-height:1.7}
.promo-code{display:inline-flex;align-items:center;gap:10px;margin-top:18px;padding:10px 18px;border-radius:10px;background:rgba(255,255,255,.06);border:1px dashed rgba(200,164,90,.4)}
.pcode{font-family:monospace;font-size:1.1rem;font-weight:700;color:var(--gold);letter-spacing:2px}
/* SHOP */
.shop-layout{max-width:1280px;margin:0 auto;padding:36px 36px 80px;display:grid;grid-template-columns:250px 1fr;gap:28px}
.sidebar{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:26px;height:fit-content;position:sticky;top:calc(var(--nav)+18px)}
.stitle2{font-family:'Playfair Display',serif;font-size:1.05rem;font-weight:700;margin-bottom:22px;padding-bottom:14px;border-bottom:1px solid var(--border)}
.fsec{margin-bottom:26px}
.flbl{font-size:.7rem;font-weight:700;letter-spacing:1.5px;text-transform:uppercase;color:var(--sub);margin-bottom:11px}
.fopts{display:flex;flex-direction:column;gap:5px}
.fopt{display:flex;align-items:center;gap:9px;cursor:pointer;padding:5px 8px;border-radius:8px;transition:.15s}
.fopt:hover{background:rgba(255,255,255,.05)}
.fopt input{width:15px;height:15px;accent-color:var(--gold);cursor:pointer}
.fopt label{font-size:.83rem;cursor:pointer;flex:1}
.fopt .cnt{font-size:.7rem;color:var(--muted)}
.pinputs{display:flex;gap:8px;align-items:center}
.pinputs input{flex:1;background:var(--panel);border:1px solid var(--border2);border-radius:8px;padding:7px 10px;color:var(--text);font-size:.83rem;outline:none;font-family:'Outfit',sans-serif}
.pinputs input:focus{border-color:var(--gold)}
.psep{color:var(--muted);font-size:.8rem}
.rslider{width:100%;margin-top:8px;accent-color:var(--gold)}
.topbar{display:flex;align-items:center;justify-content:space-between;margin-bottom:20px;flex-wrap:wrap;gap:10px}
.pcount{font-size:.86rem;color:var(--sub)}
.pcount strong{color:var(--text)}
.ssort{background:var(--card);border:1px solid var(--border2);color:var(--text);padding:8px 13px;border-radius:10px;font-size:.83rem;outline:none;font-family:'Outfit',sans-serif;cursor:pointer}
.vtoggle{display:flex;gap:3px}
.vbtn{width:33px;height:33px;border-radius:8px;border:1px solid var(--border);background:transparent;display:flex;align-items:center;justify-content:center;font-size:.83rem;color:var(--sub);transition:.2s}
.vbtn.active,.vbtn:hover{background:var(--card);border-color:var(--gold);color:var(--gold)}
.sbar{display:flex;align-items:center;gap:10px;background:var(--card);border:1px solid var(--border2);border-radius:12px;padding:0 14px;margin-bottom:18px}
.sbar input{flex:1;background:transparent;border:none;outline:none;color:var(--text);font-size:.88rem;padding:11px 0;font-family:'Outfit',sans-serif}
.sbar input::placeholder{color:var(--muted)}
.pgrid2{display:grid;grid-template-columns:repeat(auto-fill,minmax(210px,1fr));gap:18px}
.pgrid2.lv{grid-template-columns:1fr}
.pgrid2.lv .pcard{display:flex;flex-direction:row}
.pgrid2.lv .pimg{width:150px;flex-shrink:0;aspect-ratio:1}
.pgrid2.lv .pbody{flex:1}
.ftoggle{display:none;width:100%;margin-bottom:14px}
.sbclose{display:none}
/* DETAIL */
.dw{max-width:1280px;margin:0 auto;padding:36px 36px 80px}
.bc{display:flex;align-items:center;gap:8px;font-size:.8rem;color:var(--sub);margin-bottom:28px}
.bc span{color:var(--muted)}.bc a:hover{color:var(--gold)}
.dgrid{display:grid;grid-template-columns:1fr 1fr;gap:56px;align-items:start}
.dimgs{position:sticky;top:calc(var(--nav)+18px)}
.dmain{width:100%;aspect-ratio:1;object-fit:cover;border-radius:var(--r2);background:var(--card);border:1px solid var(--border)}
.dthumbs{display:flex;gap:9px;margin-top:10px;overflow-x:auto;padding-bottom:4px}
.dthumb{width:66px;height:66px;object-fit:cover;border-radius:10px;border:2px solid var(--border);cursor:pointer;flex-shrink:0;background:var(--card);transition:.2s}
.dthumb:hover,.dthumb.active{border-color:var(--gold)}
.dbadge{display:inline-block;font-size:.68rem;font-weight:700;letter-spacing:1px;text-transform:uppercase;padding:3px 12px;border-radius:6px;margin-bottom:14px;background:rgba(200,164,90,.1);color:var(--gold);border:1px solid rgba(200,164,90,.25)}
.dtitle{font-family:'Playfair Display',serif;font-size:clamp(1.5rem,3vw,2.3rem);font-weight:700;line-height:1.2;margin-bottom:10px}
.drating{display:flex;align-items:center;gap:10px;margin-bottom:18px}
.stars{color:var(--gold);letter-spacing:1px}
.rcount{font-size:.8rem;color:var(--sub)}
.dprice-w{display:flex;align-items:center;gap:14px;margin-bottom:22px}
.dprice{font-family:'Playfair Display',serif;font-size:2.1rem;font-weight:700;color:var(--gold)}
.dorig{font-size:1.05rem;color:var(--muted);text-decoration:line-through}
.ddisc{font-size:.8rem;font-weight:700;color:var(--red);background:rgba(255,77,112,.1);padding:3px 10px;border-radius:6px}
.ddesc{color:var(--sub);line-height:1.8;font-size:.93rem;margin-bottom:26px}
.dspecs{display:grid;grid-template-columns:1fr 1fr;gap:9px;margin-bottom:26px}
.specit{background:var(--card);border:1px solid var(--border);border-radius:10px;padding:11px 13px}
.speck{font-size:.7rem;color:var(--muted);letter-spacing:.5px;text-transform:uppercase;margin-bottom:3px}
.specv{font-size:.86rem;font-weight:600}
.dqty{display:flex;align-items:center;gap:14px;margin-bottom:22px}
.qlbl{font-size:.8rem;color:var(--sub);font-weight:500}
.qctl{display:flex;align-items:center;border:1px solid var(--border2);border-radius:10px;overflow:hidden}
.qbtn{width:37px;height:37px;background:transparent;border:none;color:var(--text);font-size:1.1rem;display:flex;align-items:center;justify-content:center;transition:.2s}
.qbtn:hover{background:rgba(255,255,255,.08)}
.qnum{width:42px;text-align:center;font-size:.93rem;font-weight:600;background:transparent;border:none;border-left:1px solid var(--border);border-right:1px solid var(--border);color:var(--text);padding:7px 0;font-family:'Outfit',sans-serif}
.dacts{display:flex;gap:11px;flex-wrap:wrap;margin-bottom:22px}
.dacts .btn{flex:1;min-width:110px;justify-content:center;display:flex;align-items:center;gap:7px}
.stockb{font-size:.78rem;padding:5px 13px;border-radius:8px}
.ins{background:rgba(34,211,160,.1);color:var(--green);border:1px solid rgba(34,211,160,.2)}
.outs{background:rgba(255,77,112,.1);color:var(--red);border:1px solid rgba(255,77,112,.2)}
.dtags{display:flex;gap:7px;flex-wrap:wrap;margin-top:18px}
.tag{font-size:.7rem;padding:4px 11px;border-radius:6px;border:1px solid var(--border);color:var(--sub);cursor:pointer;transition:.2s}
.tag:hover{border-color:var(--gold);color:var(--gold)}
/* CART */
.cw{max-width:1100px;margin:0 auto;padding:36px 36px 80px}
.ctitle{font-family:'Playfair Display',serif;font-size:1.9rem;font-weight:700;margin-bottom:28px}
.clayout{display:grid;grid-template-columns:1fr 360px;gap:24px;align-items:start}
.citems{display:flex;flex-direction:column;gap:14px}
.citem{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:18px;display:flex;gap:18px;align-items:center;transition:.25s}
.citem:hover{border-color:var(--border2)}
.cimg{width:86px;height:86px;object-fit:cover;border-radius:12px;background:var(--panel);flex-shrink:0;cursor:pointer}
.cinfo{flex:1}
.cname{font-weight:600;font-size:.93rem;margin-bottom:3px;cursor:pointer}
.cname:hover{color:var(--gold)}
.ccat{font-size:.75rem;color:var(--muted);margin-bottom:8px;text-transform:uppercase;letter-spacing:.5px}
.cprice{font-family:'Playfair Display',serif;font-size:1.05rem;font-weight:700;color:var(--gold)}
.cacts{display:flex;align-items:center;gap:14px;flex-shrink:0}
.rmbtn{width:31px;height:31px;border-radius:8px;border:1px solid rgba(255,77,112,.3);background:rgba(255,77,112,.08);color:var(--red);display:flex;align-items:center;justify-content:center;font-size:.88rem;transition:.2s}
.rmbtn:hover{background:var(--red);color:#fff;border-color:var(--red)}
.cempty{text-align:center;padding:70px 20px;color:var(--sub)}
.cempty .ei{font-size:3.5rem;margin-bottom:18px;opacity:.4}
.csum{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:26px;position:sticky;top:calc(var(--nav)+18px)}
.sumtitle{font-family:'Playfair Display',serif;font-size:1.15rem;font-weight:700;margin-bottom:22px;padding-bottom:14px;border-bottom:1px solid var(--border)}
.srow{display:flex;justify-content:space-between;align-items:center;margin-bottom:12px;font-size:.88rem}
.srow.tot{font-size:1.05rem;font-weight:700;border-top:1px solid var(--border);padding-top:14px;margin-top:6px}
.srow .slbl{color:var(--sub)}.srow .sval{font-weight:600}
.srow.tot .sval{color:var(--gold);font-family:'Playfair Display',serif;font-size:1.35rem}
.crow{display:flex;gap:7px;margin:18px 0}
.crow input{flex:1;background:var(--panel);border:1px solid var(--border2);border-radius:9px;padding:9px 13px;color:var(--text);font-size:.83rem;outline:none;font-family:'Outfit',sans-serif}
.crow input:focus{border-color:var(--gold)}
.freeship{display:flex;align-items:center;gap:7px;font-size:.78rem;color:var(--green);background:rgba(34,211,160,.08);padding:9px 13px;border-radius:8px;border:1px solid rgba(34,211,160,.15);margin-bottom:14px}
/* PAYMENT */
.pw{max-width:980px;margin:0 auto;padding:36px 36px 80px}
.ptitle{font-family:'Playfair Display',serif;font-size:1.9rem;font-weight:700;margin-bottom:7px}
.psub{color:var(--sub);margin-bottom:32px;font-size:.93rem}
.playout{display:grid;grid-template-columns:1fr 340px;gap:24px;align-items:start}
.pform{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:28px}
.steps{display:flex;align-items:center;gap:0;margin-bottom:28px}
.step{display:flex;align-items:center;gap:7px}
.snum{width:30px;height:30px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-size:.78rem;font-weight:700;flex-shrink:0}
.step.done .snum{background:var(--green);color:#000}
.step.active .snum{background:var(--gold);color:#000}
.step.off .snum{background:var(--panel);border:1px solid var(--border);color:var(--muted)}
.sname{font-size:.8rem;font-weight:500}
.step.off .sname{color:var(--muted)}
.sline{flex:1;height:1px;background:var(--border);margin:0 10px}
.sline.done{background:var(--green)}
.fsec2{margin-bottom:26px}
.fstitle{font-size:.72rem;font-weight:700;letter-spacing:1.5px;text-transform:uppercase;color:var(--sub);margin-bottom:14px;padding-bottom:9px;border-bottom:1px solid var(--border)}
.frow{display:grid;grid-template-columns:1fr 1fr;gap:14px}
.fg{margin-bottom:14px}
.fg label{display:block;font-size:.76rem;font-weight:500;color:var(--sub);margin-bottom:5px}
.fg input,.fg select{width:100%;background:var(--panel);border:1px solid var(--border2);border-radius:9px;padding:10px 13px;color:var(--text);font-size:.86rem;outline:none;font-family:'Outfit',sans-serif;transition:.2s}
.fg input:focus,.fg select:focus{border-color:var(--gold);box-shadow:0 0 0 3px rgba(200,164,90,.1)}
.fg select option{background:var(--panel)}
.pmethods{display:flex;flex-direction:column;gap:9px}
.pmeth{display:flex;align-items:center;gap:12px;padding:13px 15px;border-radius:12px;border:1.5px solid var(--border);cursor:pointer;transition:.2s}
.pmeth:hover{border-color:var(--border2)}
.pmeth.sel{border-color:var(--gold);background:rgba(200,164,90,.05)}
.pmeth input{accent-color:var(--gold)}
.pmico{font-size:1.35rem;width:34px;text-align:center}
.pminfo .pmname{font-weight:600;font-size:.88rem}
.pminfo .pmdesc{font-size:.73rem;color:var(--sub)}
.cfields{padding:14px;background:var(--panel);border-radius:10px;margin-top:7px;border:1px solid var(--border)}
.cicons{display:flex;gap:7px;margin-top:10px}
.cion{padding:3px 9px;border-radius:5px;background:rgba(255,255,255,.08);font-size:.7rem;font-weight:700;letter-spacing:.5px}
.osum{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:22px}
.oit{display:flex;gap:11px;align-items:center;margin-bottom:12px;padding-bottom:12px;border-bottom:1px solid var(--border)}
.oit:last-child{border-bottom:none;margin-bottom:0;padding-bottom:0}
.oimg{width:52px;height:52px;object-fit:cover;border-radius:8px;background:var(--panel);flex-shrink:0}
.oname{font-size:.83rem;font-weight:600;margin-bottom:1px}
.oprice{font-size:.83rem;color:var(--gold);font-weight:700}
.sov{display:none;position:fixed;inset:0;z-index:1000;background:rgba(0,0,0,.82);backdrop-filter:blur(10px);align-items:center;justify-content:center}
.sov.show{display:flex}
.scard{background:var(--card);border:1px solid rgba(34,211,160,.3);border-radius:var(--r2);padding:44px;text-align:center;max-width:420px;width:90%;animation:popIn .4s cubic-bezier(.34,1.56,.64,1)}
@keyframes popIn{0%{transform:scale(.8);opacity:0}100%{transform:scale(1);opacity:1}}
.sico{font-size:3.8rem;margin-bottom:18px}
.scard h2{font-family:'Playfair Display',serif;font-size:1.75rem;margin-bottom:8px}
.scard p{color:var(--sub);line-height:1.7;margin-bottom:20px}
/* GALLERY */
.gw{max-width:1280px;margin:0 auto;padding:36px 36px 80px}
.ghead{margin-bottom:44px}
.gfilts{display:flex;gap:7px;flex-wrap:wrap;margin-top:20px}
.gfb{padding:7px 18px;border-radius:99px;border:1px solid var(--border);background:transparent;color:var(--sub);font-size:.81rem;font-weight:500;transition:.2s;cursor:pointer;font-family:'Outfit',sans-serif}
.gfb.active,.gfb:hover{background:var(--gold);border-color:var(--gold);color:#000;font-weight:600}
.masonry{columns:4;column-gap:14px}
.mitem{break-inside:avoid;margin-bottom:14px;border-radius:var(--r);overflow:hidden;cursor:pointer;position:relative;background:var(--card)}
.mitem img{width:100%;display:block;transition:transform .4s}
.mitem:hover img{transform:scale(1.06)}
.movl{position:absolute;inset:0;background:linear-gradient(to top,rgba(0,0,0,.7),transparent);opacity:0;transition:.3s;display:flex;align-items:flex-end;padding:14px}
.mitem:hover .movl{opacity:1}
.mlbl{font-size:.78rem;font-weight:600}
.lbox{display:none;position:fixed;inset:0;z-index:900;background:rgba(0,0,0,.93);backdrop-filter:blur(20px);align-items:center;justify-content:center}
.lbox.open{display:flex}
.lbox-img{max-width:90vw;max-height:85vh;object-fit:contain;border-radius:16px}
.lbox-close{position:absolute;top:22px;right:22px;width:40px;height:40px;border-radius:50%;background:rgba(255,255,255,.1);border:1px solid rgba(255,255,255,.2);color:#fff;font-size:1.15rem;display:flex;align-items:center;justify-content:center;cursor:pointer;transition:.2s}
.lbox-close:hover{background:rgba(255,255,255,.2)}
/* ABOUT */
.aw{max-width:1100px;margin:0 auto;padding:56px 36px 80px}
.ahero{display:grid;grid-template-columns:1fr 1fr;gap:70px;align-items:center;margin-bottom:90px}
.aimg-w{position:relative}
.aimg{width:100%;aspect-ratio:4/5;object-fit:cover;border-radius:var(--r2);background:var(--card)}
.aaccent{position:absolute;width:180px;height:180px;border-radius:var(--r2);background:linear-gradient(135deg,rgba(124,92,245,.25),rgba(0,200,176,.18));border:1px solid rgba(124,92,245,.3);bottom:-22px;right:-22px;z-index:-1}
.vgrid{display:grid;grid-template-columns:repeat(3,1fr);gap:18px;margin-bottom:90px}
.vcard{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);padding:26px;transition:.3s}
.vcard:hover{border-color:rgba(200,164,90,.3);transform:translateY(-4px)}
.vicon{font-size:1.7rem;margin-bottom:14px}
.vtitle{font-weight:700;margin-bottom:7px}
.vdesc{font-size:.83rem;color:var(--sub);line-height:1.7}
.tgrid{display:grid;grid-template-columns:repeat(auto-fill,minmax(210px,1fr));gap:22px}
.tcard{background:var(--card);border:1px solid var(--border);border-radius:var(--r2);overflow:hidden;text-align:center;padding-bottom:22px;transition:.3s}
.tcard:hover{border-color:rgba(200,164,90,.2);transform:translateY(-4px)}
.timg{width:100%;aspect-ratio:1;object-fit:cover;background:var(--panel)}
.tname{font-weight:700;margin:14px 0 3px;font-size:.93rem}
.trole{font-size:.76rem;color:var(--gold);letter-spacing:.5px}
/* TOAST */
.toast{position:fixed;bottom:26px;right:26px;z-index:999;background:var(--panel);border:1px solid var(--border2);border-radius:12px;padding:13px 18px;display:flex;align-items:center;gap:11px;box-shadow:0 8px 32px rgba(0,0,0,.4);transform:translateY(100px);opacity:0;transition:all .35s cubic-bezier(.34,1.4,.64,1);pointer-events:none;max-width:310px}
.toast.show{transform:translateY(0);opacity:1;pointer-events:auto}
.ticon{font-size:1.15rem}
.tmsg{font-size:.86rem;font-weight:500}
/* FOOTER */
footer{background:var(--deep);border-top:1px solid var(--border);padding:56px 36px 28px}
.fin{max-width:1280px;margin:0 auto}
.fgrid{display:grid;grid-template-columns:2fr 1fr 1fr 1fr;gap:44px;margin-bottom:44px}
.fbrand .nav-logo{margin-bottom:14px;font-size:1.55rem}
.fbrand p{color:var(--sub);font-size:.86rem;line-height:1.7;max-width:270px}
.fsoc{display:flex;gap:9px;margin-top:18px}
.sbtn{width:35px;height:35px;border-radius:9px;border:1px solid var(--border);background:transparent;display:flex;align-items:center;justify-content:center;font-size:.88rem;cursor:pointer;transition:.2s;color:var(--sub)}
.sbtn:hover{background:var(--card);border-color:var(--gold);color:var(--gold)}
.fcol h4{font-size:.73rem;font-weight:700;letter-spacing:1.5px;text-transform:uppercase;color:var(--sub);margin-bottom:14px}
.fcol ul{list-style:none}
.fcol li{margin-bottom:8px}
.fcol a{font-size:.83rem;color:var(--muted);transition:.15s}
.fcol a:hover{color:var(--gold)}
.fbot{display:flex;justify-content:space-between;align-items:center;border-top:1px solid var(--border);padding-top:22px;font-size:.76rem;color:var(--muted);flex-wrap:wrap;gap:10px}
/* SKELETON */
.skel{background:linear-gradient(90deg,var(--card) 25%,var(--panel) 50%,var(--card) 75%);background-size:200% 100%;animation:shim 1.4s infinite;border-radius:10px}
@keyframes shim{0%{background-position:200% 0}100%{background-position:-200% 0}}
.hidden{display:none!important}
/* RESPONSIVE */
@media(max-width:1024px){
  .hero-vis{width:42vw;right:-16px}
  .shop-layout{grid-template-columns:1fr}
  .sidebar{position:fixed;left:0;top:0;bottom:0;z-index:600;transform:translateX(-100%);transition:.3s;border-radius:0;height:100vh;overflow-y:auto;padding-top:calc(var(--nav)+18px)}
  .sidebar.open{transform:translateX(0)}
  .ftoggle{display:block}
  .sbclose{display:flex}
  .dgrid{grid-template-columns:1fr}
  .dimgs{position:static}
  .clayout{grid-template-columns:1fr}
  .playout{grid-template-columns:1fr}
  .masonry{columns:3}
  .ahero{grid-template-columns:1fr;gap:36px}
  .vgrid{grid-template-columns:1fr 1fr}
  .fgrid{grid-template-columns:1fr 1fr}
}
@media(max-width:768px){
  nav{padding:0 18px}
  .nav-links{display:none}
  .ham{display:flex}
  .hero{padding:56px 18px 36px;min-height:auto}
  .hero-vis{display:none}
  .hero-stats{gap:20px;flex-wrap:wrap}
  .featured,.categories,.promo{padding-left:18px;padding-right:18px}
  .promo-in{flex-direction:column;padding:32px 22px}
  .shop-layout,.dw,.cw,.pw,.gw,.aw{padding-left:16px;padding-right:16px}
  .masonry{columns:2}
  .frow{grid-template-columns:1fr}
  .fgrid{grid-template-columns:1fr}
  .citem{flex-wrap:wrap}
  .dspecs{grid-template-columns:1fr}
  .vgrid{grid-template-columns:1fr}
  .ahero{margin-bottom:56px}
}
@media(max-width:480px){
  .masonry{columns:1}
  .hero h1{font-size:2.1rem}
  .dacts{flex-direction:column}
  .cacts{width:100%;justify-content:space-between}
  .pgrid,.pgrid2{grid-template-columns:repeat(auto-fill,minmax(160px,1fr))}
}
</style>
</head>
<body>
<div id="amb"><span></span><span></span><span></span></div>

<nav>
  <div class="nav-logo">ASLAN</div>
  <div class="nav-links">
    <a href="#" class="active" onclick="go('home',this)">Home</a>
    <a href="#" onclick="go('shop',this)">Shop</a>
    <a href="#" onclick="go('gallery',this)">Gallery</a>
    <a href="#" onclick="go('about',this)">About</a>
  </div>
  <div class="nav-right">
    <button class="cart-btn" onclick="go('cart',null)">🛒<span class="cart-ct" id="cct">0</span></button>
    <div class="ham" onclick="toggleMM()"><span></span><span></span><span></span></div>
  </div>
</nav>
<div class="mmenu" id="mmenu">
  <a href="#" onclick="go('home',this);toggleMM()">🏠 Home</a>
  <a href="#" onclick="go('shop',this);toggleMM()">🛍 Shop</a>
  <a href="#" onclick="go('gallery',this);toggleMM()">🖼 Gallery</a>
  <a href="#" onclick="go('about',this);toggleMM()">ℹ About</a>
  <a href="#" onclick="go('cart',null);toggleMM()">🛒 Cart (<span id="mcct">0</span>)</a>
</div>
<div class="toast" id="toast"><span class="ticon" id="tico">✓</span><span class="tmsg" id="tmsg"></span></div>
<div class="sov" id="sov">
  <div class="scard">
    <div class="sico">🎉</div>
    <h2>Order Placed!</h2>
    <p>Terima kasih telah berbelanja di ASLAN. Pesanan Anda sedang diproses.</p>
    <p style="font-size:.8rem;color:var(--muted)">Order ID: <strong id="oid" style="color:var(--gold)"></strong></p>
    <div style="display:flex;gap:10px;justify-content:center;margin-top:6px">
      <button class="btn btn-ghost" onclick="document.getElementById('sov').classList.remove('show');go('shop',null)">Lanjut Belanja</button>
      <button class="btn btn-gold" onclick="document.getElementById('sov').classList.remove('show');go('home',null)">Home</button>
    </div>
  </div>
</div>

<!-- HOME -->
<div id="page-home" class="page active">
  <section class="hero">
    <div class="hero-content">
      <div class="hero-eyebrow"><span>✦ NEW COLLECTION 2025</span></div>
      <h1>Discover <em>Premium</em><br>Lifestyle Goods</h1>
      <p class="hero-desc">Temukan koleksi eksklusif produk premium pilihan terbaik. Kualitas tanpa kompromi, gaya tanpa batas.</p>
      <div class="hero-actions">
        <button class="btn btn-gold" onclick="go('shop',null)">Shop Now →</button>
        <button class="btn btn-outline" onclick="go('gallery',null)">View Gallery</button>
      </div>
      <div class="hero-stats">
        <div class="hstat"><div class="num">2.4K+</div><div class="lbl">Products</div></div>
        <div class="hstat"><div class="num">98%</div><div class="lbl">Satisfied</div></div>
        <div class="hstat"><div class="num">50+</div><div class="lbl">Brands</div></div>
        <div class="hstat"><div class="num">4.9★</div><div class="lbl">Rating</div></div>
      </div>
    </div>
    <div class="hero-vis">
      <div style="position:relative">
        <img class="hero-img" src="https://images.unsplash.com/photo-1441984904996-e0b6ba687e04?w=600&q=80" alt="hero"/>
        <div class="hfc"><div class="fl-lbl">⭐ TOP RATED</div><div class="fl-val">4.9 / 5.0</div></div>
        <div class="hfc"><div class="fl-lbl">🚚 FREE SHIPPING</div><div class="fl-val">Order > $50</div></div>
      </div>
    </div>
  </section>
  <div class="mq">
    <div class="mq-track">
      <span class="mq-item"><span>✦</span>FREE SHIPPING ABOVE $50</span>
      <span class="mq-item"><span>✦</span>NEW ARRIVALS EVERY WEEK</span>
      <span class="mq-item"><span>✦</span>PREMIUM QUALITY GUARANTEED</span>
      <span class="mq-item"><span>✦</span>30-DAY EASY RETURNS</span>
      <span class="mq-item"><span>✦</span>EXCLUSIVE MEMBER DISCOUNTS</span>
      <span class="mq-item"><span>✦</span>SECURE PAYMENTS</span>
      <span class="mq-item"><span>✦</span>FREE SHIPPING ABOVE $50</span>
      <span class="mq-item"><span>✦</span>NEW ARRIVALS EVERY WEEK</span>
      <span class="mq-item"><span>✦</span>PREMIUM QUALITY GUARANTEED</span>
      <span class="mq-item"><span>✦</span>30-DAY EASY RETURNS</span>
      <span class="mq-item"><span>✦</span>EXCLUSIVE MEMBER DISCOUNTS</span>
      <span class="mq-item"><span>✦</span>SECURE PAYMENTS</span>
    </div>
  </div>
  <section class="categories">
    <div style="text-align:center"><div class="stag">Browse By</div><h2 class="stitle">Shop Categories</h2></div>
    <div class="cgrid" id="cgrid"></div>
  </section>
  <section class="featured">
    <div style="display:flex;justify-content:space-between;align-items:flex-end;flex-wrap:wrap;gap:14px">
      <div><div class="stag">Featured</div><h2 class="stitle">Best Sellers</h2><p class="ssub">Produk paling dicari dengan rating terbaik dari pelanggan kami.</p></div>
      <button class="btn btn-outline" onclick="go('shop',null)">View All →</button>
    </div>
    <div class="pgrid" id="feat-grid">
      <div style="height:280px" class="skel"></div><div style="height:280px" class="skel"></div>
      <div style="height:280px" class="skel"></div><div style="height:280px" class="skel"></div>
    </div>
  </section>
  <section class="promo">
    <div class="promo-in">
      <div class="promo-txt">
        <div class="stag">Limited Offer</div>
        <h2>Get 20% Off Your<br>First Purchase</h2>
        <p>Daftar sekarang dan dapatkan diskon eksklusif untuk pembelian pertama Anda di ASLAN Store.</p>
        <div class="promo-code"><span style="font-size:.76rem;color:var(--sub)">Use code:</span><span class="pcode">ASLAN20</span><button class="btn btn-sm btn-ghost" onclick="copyCode()">Copy</button></div>
      </div>
      <button class="btn btn-gold" style="flex-shrink:0" onclick="go('shop',null)">Claim Discount →</button>
    </div>
  </section>
</div>

<!-- SHOP -->
<div id="page-shop" class="page">
  <div class="shop-layout">
    <aside class="sidebar" id="sidebar">
      <button class="btn btn-ghost sbclose" onclick="closeSB()">✕ Close</button>
      <div class="stitle2">🔍 Filters</div>
      <div class="fsec"><div class="flbl">Category</div><div class="fopts" id="cat-filters"></div></div>
      <div class="fsec">
        <div class="flbl">Price Range</div>
        <div class="pinputs">
          <input type="number" id="pmin" placeholder="Min" value="0" oninput="af()"/>
          <span class="psep">—</span>
          <input type="number" id="pmax" placeholder="Max" value="2000" oninput="af()"/>
        </div>
        <input type="range" class="rslider" min="0" max="2000" value="2000" id="prange" oninput="document.getElementById('pmax').value=this.value;af()"/>
      </div>
      <div class="fsec">
        <div class="flbl">Min Rating</div>
        <div class="fopts">
          <label class="fopt"><input type="radio" name="rat" value="0" checked onchange="af()"/><label>All</label></label>
          <label class="fopt"><input type="radio" name="rat" value="4" onchange="af()"/><label>⭐ 4+</label></label>
          <label class="fopt"><input type="radio" name="rat" value="4.5" onchange="af()"/><label>⭐ 4.5+</label></label>
        </div>
      </div>
      <button class="btn btn-gold" style="width:100%;margin-top:6px" onclick="rf()">Reset Filters</button>
    </aside>
    <div>
      <button class="btn btn-ghost ftoggle" onclick="openSB()">⚙ Filter & Sort</button>
      <div class="sbar"><span style="color:var(--muted)">🔍</span><input type="text" id="ssearch" placeholder="Search products..." oninput="af()"/></div>
      <div class="topbar">
        <div class="pcount">Showing <strong id="pcnt">0</strong> products</div>
        <div style="display:flex;gap:9px;align-items:center">
          <select class="ssort" id="ssort" onchange="af()">
            <option value="default">Default</option>
            <option value="price-asc">Price: Low→High</option>
            <option value="price-desc">Price: High→Low</option>
            <option value="rating">Best Rating</option>
            <option value="name">Name A–Z</option>
          </select>
          <div class="vtoggle">
            <button class="vbtn active" id="gbtn" onclick="sv('grid')" title="Grid">⊞</button>
            <button class="vbtn" id="lbtn" onclick="sv('list')" title="List">☰</button>
          </div>
        </div>
      </div>
      <div class="pgrid2" id="pgrid2">
        <div style="height:250px" class="skel"></div><div style="height:250px" class="skel"></div>
        <div style="height:250px" class="skel"></div><div style="height:250px" class="skel"></div>
      </div>
      <div style="text-align:center;margin-top:36px"><button class="btn btn-outline" id="lmbtn" onclick="lm()" style="display:none">Load More</button></div>
    </div>
  </div>
</div>

<!-- DETAIL -->
<div id="page-detail" class="page">
  <div class="dw">
    <div class="bc">
      <a href="#" onclick="go('home',null)">Home</a><span>/</span>
      <a href="#" onclick="go('shop',null)">Shop</a><span>/</span>
      <span id="dbread">Product</span>
    </div>
    <div class="dgrid">
      <div class="dimgs">
        <img class="dmain" id="dmain" src="" alt=""/>
        <div class="dthumbs" id="dthumbs"></div>
      </div>
      <div>
        <div class="dbadge" id="dbadge">Premium</div>
        <h1 class="dtitle" id="dtitle">Loading...</h1>
        <div class="drating">
          <div class="stars" id="dstars">★★★★★</div>
          <span class="rcount" id="drev">0 reviews</span>
          <span class="stockb ins" id="dstock">In Stock</span>
        </div>
        <div class="dprice-w">
          <div class="dprice" id="dprice">$0</div>
          <div class="dorig hidden" id="dorig">$0</div>
          <div class="ddisc hidden" id="ddisc">0% OFF</div>
        </div>
        <p class="ddesc" id="ddesc">Loading...</p>
        <div class="dspecs" id="dspecs"></div>
        <div class="dqty">
          <span class="qlbl">Qty:</span>
          <div class="qctl">
            <button class="qbtn" onclick="cq(-1)">−</button>
            <input class="qnum" id="dqty" type="number" value="1" min="1" max="99"/>
            <button class="qbtn" onclick="cq(1)">+</button>
          </div>
        </div>
        <div class="dacts">
          <button class="btn btn-gold" onclick="addDTC()">🛒 Add to Cart</button>
          <button class="btn btn-teal" onclick="buyNow()">⚡ Buy Now</button>
          <button class="btn btn-outline btn-sm" onclick="wl()" id="wbtn">♡</button>
        </div>
        <div class="dtags" id="dtags"></div>
      </div>
    </div>
  </div>
</div>

<!-- CART -->
<div id="page-cart" class="page">
  <div class="cw">
    <div style="display:flex;align-items:center;justify-content:space-between;margin-bottom:28px">
      <h1 class="ctitle" style="margin-bottom:0">Shopping Cart</h1>
      <button class="btn btn-ghost btn-sm" onclick="go('shop',null)">← Continue Shopping</button>
    </div>
    <div id="cart-content"></div>
  </div>
</div>

<!-- PAYMENT -->
<div id="page-payment" class="page">
  <div class="pw">
    <h1 class="ptitle">Checkout</h1>
    <p class="psub">Lengkapi informasi pengiriman dan pembayaran Anda</p>
    <div class="playout">
      <div class="pform">
        <div class="steps">
          <div class="step done"><div class="snum">✓</div><div class="sname">Cart</div></div>
          <div class="sline done"></div>
          <div class="step active"><div class="snum">2</div><div class="sname">Details</div></div>
          <div class="sline"></div>
          <div class="step off"><div class="snum">3</div><div class="sname">Confirm</div></div>
        </div>
        <div class="fsec2">
          <div class="fstitle">📦 Shipping Info</div>
          <div class="frow">
            <div class="fg"><label>First Name</label><input id="fn" placeholder="John"/></div>
            <div class="fg"><label>Last Name</label><input id="ln" placeholder="Doe"/></div>
          </div>
          <div class="fg"><label>Email</label><input id="em" type="email" placeholder="john@example.com"/></div>
          <div class="fg"><label>Phone</label><input id="ph" type="tel" placeholder="+62 812 3456 7890"/></div>
          <div class="fg"><label>Address</label><input id="ad" placeholder="Jl. Sudirman No. 1"/></div>
          <div class="frow">
            <div class="fg"><label>City</label><input id="ct" placeholder="Jakarta"/></div>
            <div class="fg"><label>Postal Code</label><input id="zp" placeholder="12345"/></div>
          </div>
          <div class="fg"><label>Province</label>
            <select id="pv"><option>DKI Jakarta</option><option>Jawa Barat</option><option>Jawa Tengah</option><option>Jawa Timur</option><option>Bali</option><option>Sumatera Utara</option><option>Kalimantan Timur</option><option>Sulawesi Selatan</option></select>
          </div>
        </div>
        <div class="fsec2">
          <div class="fstitle">💳 Payment Method</div>
          <div class="pmethods">
            <label class="pmeth sel" id="pm-card"><input type="radio" name="pay" value="card" checked onchange="tpf()"/><span class="pmico">💳</span><div class="pminfo"><div class="pmname">Credit / Debit Card</div><div class="pmdesc">Visa, Mastercard, JCB</div></div></label>
            <div class="cfields" id="cfields">
              <div class="fg"><label>Card Number</label><input id="cn" placeholder="1234 5678 9012 3456" maxlength="19" oninput="fc(this)"/></div>
              <div class="frow">
                <div class="fg"><label>Expiry</label><input id="ex" placeholder="MM/YY" maxlength="5" oninput="fe(this)"/></div>
                <div class="fg"><label>CVV</label><input id="cv" placeholder="123" type="password" maxlength="4"/></div>
              </div>
              <div class="cicons"><span class="cion">VISA</span><span class="cion">MC</span><span class="cion">JCB</span></div>
            </div>
            <label class="pmeth" id="pm-transfer"><input type="radio" name="pay" value="transfer" onchange="tpf()"/><span class="pmico">🏦</span><div class="pminfo"><div class="pmname">Bank Transfer</div><div class="pmdesc">BCA, Mandiri, BNI, BRI</div></div></label>
            <label class="pmeth" id="pm-ewallet"><input type="radio" name="pay" value="ewallet" onchange="tpf()"/><span class="pmico">📱</span><div class="pminfo"><div class="pmname">E-Wallet</div><div class="pmdesc">GoPay, OVO, DANA, ShopeePay</div></div></label>
            <label class="pmeth" id="pm-cod"><input type="radio" name="pay" value="cod" onchange="tpf()"/><span class="pmico">💵</span><div class="pminfo"><div class="pmname">Cash on Delivery</div><div class="pmdesc">Bayar saat barang tiba</div></div></label>
          </div>
        </div>
        <button class="btn btn-gold" style="width:100%;padding:14px;font-size:.95rem" onclick="placeOrder()">Place Order →</button>
      </div>
      <div>
        <div class="osum">
          <div style="font-family:'Playfair Display',serif;font-size:1.05rem;font-weight:700;margin-bottom:18px;padding-bottom:12px;border-bottom:1px solid var(--border)">Order Summary</div>
          <div id="oitems"></div>
          <div style="border-top:1px solid var(--border);padding-top:14px;margin-top:2px">
            <div class="srow"><span class="slbl">Subtotal</span><span class="sval" id="psub">$0</span></div>
            <div class="srow"><span class="slbl">Shipping</span><span class="sval" style="color:var(--green)" id="pship">Free</span></div>
            <div class="srow"><span class="slbl">Tax (10%)</span><span class="sval" id="ptax">$0</span></div>
            <div class="srow tot"><span class="slbl">Total</span><span class="sval" id="ptot">$0</span></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<!-- GALLERY -->
<div id="page-gallery" class="page">
  <div class="gw">
    <div class="ghead">
      <div class="stag">Visual Showcase</div>
      <h2 class="stitle">Product Gallery</h2>
      <p class="ssub">Jelajahi koleksi visual produk premium kami yang dikurasi dengan cermat.</p>
      <div class="gfilts">
        <button class="gfb active" onclick="fg('all',this)">All</button>
        <button class="gfb" onclick="fg('electronics',this)">Electronics</button>
        <button class="gfb" onclick="fg('fashion',this)">Fashion</button>
        <button class="gfb" onclick="fg('home',this)">Home</button>
        <button class="gfb" onclick="fg('beauty',this)">Beauty</button>
      </div>
    </div>
    <div class="masonry" id="masonry"></div>
  </div>
  <div class="lbox" id="lbox" onclick="clb()">
    <button class="lbox-close" onclick="clb()">✕</button>
    <img class="lbox-img" id="lbimg" src="" alt=""/>
  </div>
</div>

<!-- ABOUT -->
<div id="page-about" class="page">
  <div class="aw">
    <div class="ahero">
      <div class="aimg-w">
        <img class="aimg" src="https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?w=600&q=80" alt="About"/>
        <div class="aaccent"></div>
      </div>
      <div>
        <div class="stag">Our Story</div>
        <h2 class="stitle">Passion for<br>Premium Quality</h2>
        <p class="ssub" style="max-width:100%">ASLAN lahir dari kecintaan mendalam terhadap produk berkualitas tinggi. Kami percaya setiap orang berhak mendapatkan yang terbaik tanpa harus mengorbankan gaya dan kenyamanan.</p>
        <p style="color:var(--sub);line-height:1.8;font-size:.93rem;margin-top:14px">Sejak 2020, kami telah melayani ribuan pelanggan di seluruh Indonesia dengan kurasi produk premium dari brand terpercaya di dunia.</p>
        <div style="display:flex;gap:14px;margin-top:26px;flex-wrap:wrap">
          <button class="btn btn-gold" onclick="go('shop',null)">Shop Now</button>
          <button class="btn btn-outline" onclick="go('gallery',null)">Our Gallery</button>
        </div>
      </div>
    </div>
    <div style="margin-bottom:56px">
      <div style="text-align:center;margin-bottom:36px"><div class="stag">Why Us</div><h2 class="stitle">Our Values</h2></div>
      <div class="vgrid">
        <div class="vcard"><div class="vicon">🏆</div><div class="vtitle">Premium Quality</div><div class="vdesc">Setiap produk melalui seleksi ketat untuk kualitas terbaik.</div></div>
        <div class="vcard"><div class="vicon">🚀</div><div class="vtitle">Fast Delivery</div><div class="vdesc">Pengiriman cepat ke seluruh Indonesia dengan packaging aman.</div></div>
        <div class="vcard"><div class="vicon">🛡</div><div class="vtitle">Secure Shopping</div><div class="vdesc">Transaksi 100% aman dengan enkripsi SSL terpercaya.</div></div>
        <div class="vcard"><div class="vicon">💬</div><div class="vtitle">24/7 Support</div><div class="vdesc">Tim customer service siap membantu kapan saja.</div></div>
        <div class="vcard"><div class="vicon">🔄</div><div class="vtitle">Easy Returns</div><div class="vdesc">30-hari garansi kepuasan. Tidak puas, uang kembali.</div></div>
        <div class="vcard"><div class="vicon">🌿</div><div class="vtitle">Sustainable</div><div class="vdesc">Komitmen pada praktik bisnis ramah lingkungan.</div></div>
      </div>
    </div>
    <div>
      <div style="text-align:center;margin-bottom:36px"><div class="stag">The Team</div><h2 class="stitle">Meet Our Team</h2></div>
      <div class="tgrid">
        <div class="tcard"><img class="timg" src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=300&q=70" alt=""/><div class="tname">Ariel Aslan</div><div class="trole">Founder & CEO</div></div>
        <div class="tcard"><img class="timg" src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=300&q=70" alt=""/><div class="tname">Sila ASLAN</div><div class="trole">Head of Design</div></div>
        <div class="tcard"><img class="timg" src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=300&q=70" alt=""/><div class="tname">Raka Mahendra</div><div class="trole">Lead Engineer</div></div>
        <div class="tcard"><img class="timg" src="https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=300&q=70" alt=""/><div class="tname">Dian Puspita</div><div class="trole">Marketing Director</div></div>
      </div>
    </div>
  </div>
</div>

<footer>
  <div class="fin">
    <div class="fgrid">
      <div class="fbrand">
        <div class="nav-logo">ASLAN</div>
        <p>Premium lifestyle store menghadirkan produk terbaik dari seluruh dunia untuk Anda.</p>
        <div class="fsoc">
          <button class="sbtn">𝕏</button><button class="sbtn">in</button>
          <button class="sbtn">📷</button><button class="sbtn">▶</button>
        </div>
      </div>
      <div class="fcol"><h4>Shop</h4><ul>
        <li><a href="#" onclick="go('shop',null)">All Products</a></li>
        <li><a href="#">New Arrivals</a></li>
        <li><a href="#">Best Sellers</a></li>
        <li><a href="#">Sale</a></li>
      </ul></div>
      <div class="fcol"><h4>Company</h4><ul>
        <li><a href="#" onclick="go('about',null)">About Us</a></li>
        <li><a href="#">Careers</a></li>
        <li><a href="#">Press</a></li>
        <li><a href="#">Partners</a></li>
      </ul></div>
      <div class="fcol"><h4>Support</h4><ul>
        <li><a href="#">Help Center</a></li>
        <li><a href="#">Returns</a></li>
        <li><a href="#">Shipping Info</a></li>
        <li><a href="#">Contact Us</a></li>
      </ul></div>
    </div>
    <div class="fbot">
      <span>© 2025 ASLAN Store. All rights reserved.</span>
      <span>Made with ♥ in Indonesia · Port :65534</span>
    </div>
  </div>
</footer>

<script>
const BASE='https://dummyjson.com';
let allP=[],filtP=[],dcnt=12,cart=JSON.parse(localStorage.getItem('asl_cart')||'[]'),curP=null,gItems=[];
const CATS=[
  {v:'smartphones',i:'📱',l:'Smartphones'},{v:'laptops',i:'💻',l:'Laptops'},
  {v:'fragrances',i:'🌸',l:'Fragrances'},{v:'skincare',i:'✨',l:'Skincare'},
  {v:'groceries',i:'🛒',l:'Groceries'},{v:'home-decoration',i:'🏠',l:'Home Decor'},
  {v:'furniture',i:'🪑',l:'Furniture'},{v:'tops',i:'👕',l:'Fashion'},
  {v:'womens-bags',i:'👜',l:'Bags'},{v:'sunglasses',i:'🕶',l:'Sunglasses'},
  {v:'automotive',i:'🚗',l:'Automotive'},{v:'sports-accessories',i:'⚽',l:'Sports'}
];

function go(id,el){
  document.querySelectorAll('.page').forEach(p=>p.classList.remove('active'));
  document.getElementById('page-'+id).classList.add('active');
  document.querySelectorAll('.nav-links a,.mmenu a').forEach(a=>a.classList.remove('active'));
  if(el)el.classList.add('active');
  window.scrollTo({top:0,behavior:'smooth'});
  if(id==='shop')initShop();
  if(id==='gallery')initGallery();
  if(id==='cart')renderCart();
  if(id==='payment')renderPayment();
}
function toggleMM(){document.getElementById('mmenu').classList.toggle('open')}
function openSB(){document.getElementById('sidebar').classList.add('open')}
function closeSB(){document.getElementById('sidebar').classList.remove('open')}

function toast(msg,ico='✓'){
  const t=document.getElementById('toast');
  document.getElementById('tmsg').textContent=msg;
  document.getElementById('tico').textContent=ico;
  t.classList.add('show');
  setTimeout(()=>t.classList.remove('show'),3000);
}

function saveCart(){localStorage.setItem('asl_cart',JSON.stringify(cart));ucc()}
function ucc(){
  const n=cart.reduce((s,i)=>s+i.qty,0);
  const el=document.getElementById('cct');
  const m=document.getElementById('mcct');
  el.textContent=n;el.style.display=n>0?'flex':'none';m.textContent=n;
}
function addToCart(p){
  const e=cart.find(i=>i.id===p.id);
  if(e)e.qty++;else cart.push({id:p.id,title:p.title,price:p.price,thumbnail:p.thumbnail,category:p.category,qty:1});
  saveCart();toast('Added: '+p.title.substring(0,28),'🛒');
}
function rmCart(id){cart=cart.filter(i=>i.id!==id);saveCart();renderCart()}
function updQC(id,d){const i=cart.find(x=>x.id===id);if(i){i.qty=Math.max(1,i.qty+d);saveCart();renderCart()}}

function renderCart(){
  const el=document.getElementById('cart-content');
  if(!cart.length){el.innerHTML='<div class="cempty"><div class="ei">🛒</div><h3 style="font-family:\'Playfair Display\',serif;font-size:1.35rem;margin-bottom:9px">Your cart is empty</h3><p style="margin-bottom:22px">Mulai belanja produk favoritmu!</p><button class="btn btn-gold" onclick="go(\'shop\',null)">Start Shopping</button></div>';return}
  const sub=cart.reduce((s,i)=>s+i.price*i.qty,0);
  const ship=sub>50?0:9.99;const tax=sub*.1;const tot=sub+ship+tax;
  el.innerHTML='<div class="clayout"><div class="citems">'+cart.map(i=>'<div class="citem"><img class="cimg" src="'+i.thumbnail+'" onerror="this.src=\'\'" onclick="openP('+i.id+')"/><div class="cinfo"><div class="cname" onclick="openP('+i.id+')">'+i.title+'</div><div class="ccat">'+i.category+'</div><div class="cprice">$'+(i.price*i.qty).toFixed(2)+' <span style="font-size:.74rem;color:var(--sub);font-family:\'Outfit\',sans-serif;font-weight:400">($'+i.price+' × '+i.qty+')</span></div></div><div class="cacts"><div class="qctl"><button class="qbtn" onclick="updQC('+i.id+',-1)">−</button><input class="qnum" value="'+i.qty+'" readonly/><button class="qbtn" onclick="updQC('+i.id+',1)">+</button></div><button class="rmbtn" onclick="rmCart('+i.id+')">✕</button></div></div>').join('')+'</div><div class="csum"><div class="sumtitle">Order Summary</div>'+(ship===0?'<div class="freeship">🎉 You qualify for free shipping!</div>':'')+'<div class="srow"><span class="slbl">Subtotal ('+cart.reduce((s,i)=>s+i.qty,0)+' items)</span><span class="sval">$'+sub.toFixed(2)+'</span></div><div class="srow"><span class="slbl">Shipping</span><span class="sval" style="color:'+(ship===0?'var(--green)':'var(--text)')+'">'+(ship===0?'Free':'$'+ship.toFixed(2))+'</span></div><div class="srow"><span class="slbl">Tax (10%)</span><span class="sval">$'+tax.toFixed(2)+'</span></div><div class="crow"><input placeholder="Promo code (ASLAN20)" id="pcin"/><button class="btn btn-ghost btn-sm" onclick="applyP()">Apply</button></div><div class="srow tot"><span class="slbl">Total</span><span class="sval">$'+tot.toFixed(2)+'</span></div><button class="btn btn-gold" style="width:100%;margin-top:14px;padding:13px" onclick="go(\'payment\',null)">Checkout →</button><button class="btn btn-ghost" style="width:100%;margin-top:7px" onclick="go(\'shop\',null)">Continue Shopping</button></div></div>';
}

function applyP(){const c=document.getElementById('pcin')?.value.trim().toUpperCase();c==='ASLAN20'?toast('20% off applied! 🎉','🎁'):toast('Invalid code','✕')}

async function loadP(){
  try{const r=await fetch(BASE+'/products?limit=100');const d=await r.json();allP=d.products;filtP=[...allP];buildCF();renderFeat();renderCats();return true}
  catch(e){return false}
}

function initShop(){
  if(!allP.length)loadP().then(()=>af());
  else af();
  buildCF();
}

function buildCF(){
  if(!allP.length)return;
  const cats=[...new Set(allP.map(p=>p.category))];
  const counts={};cats.forEach(c=>counts[c]=allP.filter(p=>p.category===c).length);
  document.getElementById('cat-filters').innerHTML='<label class="fopt"><input type="checkbox" value="all" onchange="af()" checked/><label>All Categories</label></label>'+cats.map(c=>'<label class="fopt"><input type="checkbox" value="'+c+'" onchange="af()"/><label>'+c.replace(/-/g,' ')+'</label><span class="cnt">'+counts[c]+'</span></label>').join('');
}

function af(){
  const q=(document.getElementById('ssearch')?.value||'').toLowerCase();
  const pmin=parseFloat(document.getElementById('pmin')?.value||0);
  const pmax=parseFloat(document.getElementById('pmax')?.value||9999);
  const rmin=parseFloat(document.querySelector('input[name="rat"]:checked')?.value||0);
  const srt=document.getElementById('ssort')?.value||'default';
  const cats=[...document.querySelectorAll('#cat-filters input[type=checkbox]:checked')].map(i=>i.value);
  const allC=cats.includes('all')||!cats.length;
  let f=allP.filter(p=>{
    if(q&&!p.title.toLowerCase().includes(q)&&!p.category.toLowerCase().includes(q))return false;
    if(p.price<pmin||p.price>pmax)return false;
    if(p.rating<rmin)return false;
    if(!allC&&!cats.includes(p.category))return false;
    return true;
  });
  if(srt==='price-asc')f.sort((a,b)=>a.price-b.price);
  else if(srt==='price-desc')f.sort((a,b)=>b.price-a.price);
  else if(srt==='rating')f.sort((a,b)=>b.rating-a.rating);
  else if(srt==='name')f.sort((a,b)=>a.title.localeCompare(b.title));
  filtP=f;dcnt=12;renderSG();
}

function rf(){
  document.getElementById('ssearch').value='';
  document.getElementById('pmin').value='0';
  document.getElementById('pmax').value='2000';
  document.getElementById('prange').value='2000';
  document.querySelector('input[name="rat"][value="0"]').checked=true;
  document.querySelectorAll('#cat-filters input').forEach(i=>i.checked=i.value==='all');
  document.getElementById('ssort').value='default';
  af();
}

function renderSG(){
  const g=document.getElementById('pgrid2');
  const shown=filtP.slice(0,dcnt);
  document.getElementById('pcnt').textContent=filtP.length;
  document.getElementById('lmbtn').style.display=filtP.length>dcnt?'inline-flex':'none';
  if(!shown.length){g.innerHTML='<div style="grid-column:1/-1;text-align:center;padding:56px 20px;color:var(--sub)"><div style="font-size:2.8rem;margin-bottom:14px">🔍</div><p>No products found.</p></div>';return}
  g.innerHTML=shown.map(p=>pc(p)).join('');
}

function pc(p){
  const bl=['b-new','b-sale','b-hot'];const bn=['NEW','SALE','HOT'];const bi=p.id%3;
  const t=p.title.replace(/'/g,"\\'");const th=p.thumbnail;const cat=p.category;
  return '<div class="pcard" onclick="openP('+p.id+')"><span class="pbadge '+bl[bi]+'">'+bn[bi]+'</span><img class="pimg" src="'+th+'" onerror="this.src=\'https://placehold.co/300x300/12121e/555?text=No+Img\'" loading="lazy"/><div class="pbody"><div class="pcat">'+cat+'</div><div class="pname">'+p.title+'</div><div class="pbottom"><div><div class="pprice">$'+p.price+'</div><div class="prating">⭐ '+p.rating+' · '+p.stock+' left</div></div><button class="padd" onclick="event.stopPropagation();addToCart({id:'+p.id+',title:\''+t+'\',price:'+p.price+',thumbnail:\''+th+'\',category:\''+cat+'\'})" title="Add to cart">+</button></div></div></div>';
}

function lm(){dcnt+=12;renderSG()}

function sv(v){
  document.getElementById('pgrid2').classList.toggle('lv',v==='list');
  document.getElementById('gbtn').classList.toggle('active',v==='grid');
  document.getElementById('lbtn').classList.toggle('active',v==='list');
}

function renderFeat(){
  const top=[...allP].sort((a,b)=>b.rating-a.rating).slice(0,8);
  document.getElementById('feat-grid').innerHTML=top.map(p=>pc(p)).join('');
}

function renderCats(){
  document.getElementById('cgrid').innerHTML=CATS.map(c=>'<div class="citem" onclick="fbyCat(\''+c.v+'\')"><div class="cicon">'+c.i+'</div><div class="cname">'+c.l+'</div></div>').join('');
}

function fbyCat(cat){
  go('shop',null);
  setTimeout(()=>{document.getElementById('ssearch').value=cat.replace(/-/g,' ');af()},100);
}

async function openP(id){
  go('detail',null);
  try{
    const r=await fetch(BASE+'/products/'+id);const p=await r.json();curP=p;
    document.getElementById('dbread').textContent=p.title.substring(0,28);
    document.getElementById('dmain').src=p.images?.[0]||p.thumbnail;
    document.getElementById('dtitle').textContent=p.title;
    document.getElementById('dprice').textContent='$'+p.price;
    document.getElementById('ddesc').textContent=p.description;
    document.getElementById('dbadge').textContent=p.category;
    document.getElementById('dstars').textContent='★'.repeat(Math.round(p.rating||0))+'☆'.repeat(5-Math.round(p.rating||0));
    document.getElementById('drev').textContent=(Math.floor(Math.random()*280+40))+' reviews';
    const si=document.getElementById('dstock');si.className='stockb '+(p.stock>0?'ins':'outs');si.textContent=p.stock>0?'In Stock ('+p.stock+')':'Out of Stock';
    const orig=(p.price*(1+(p.discountPercentage||0)/100)).toFixed(2);
    if(p.discountPercentage>0){document.getElementById('dorig').textContent='$'+orig;document.getElementById('dorig').classList.remove('hidden');document.getElementById('ddisc').textContent=Math.round(p.discountPercentage)+'% OFF';document.getElementById('ddisc').classList.remove('hidden');}
    else{document.getElementById('dorig').classList.add('hidden');document.getElementById('ddisc').classList.add('hidden')}
    const imgs=p.images||[p.thumbnail];
    document.getElementById('dthumbs').innerHTML=imgs.map((img,i)=>'<img class="dthumb'+(i===0?' active':'')+'" src="'+img+'" loading="lazy" onclick="si2(this,\''+img+'\')"/>').join('');
    document.getElementById('dspecs').innerHTML=[['Brand',p.brand||'—'],['Category',p.category||'—'],['Rating',(p.rating||0)+'/5'],['Stock',p.stock+' units'],['SKU',p.sku||'SKU-'+p.id],['Weight',(p.weight||1)+'kg']].map(([k,v])=>'<div class="specit"><div class="speck">'+k+'</div><div class="specv">'+v+'</div></div>').join('');
    document.getElementById('dtags').innerHTML=(p.tags||[p.category]).map(t=>'<span class="tag" onclick="fbyCat(\''+t+'\')">'+t+'</span>').join('');
    document.getElementById('dqty').value=1;
  }catch(e){toast('Failed to load','✕')}
}

function si2(thumb,src){
  document.getElementById('dmain').src=src;
  document.querySelectorAll('.dthumb').forEach(t=>t.classList.remove('active'));
  thumb.classList.add('active');
}

function cq(d){const i=document.getElementById('dqty');i.value=Math.max(1,parseInt(i.value||1)+d)}
function addDTC(){if(!curP)return;const q=parseInt(document.getElementById('dqty').value)||1;for(let i=0;i<q;i++)addToCart(curP)}
function buyNow(){addDTC();go('cart',null)}
function wl(){const b=document.getElementById('wbtn');const w=b.textContent==='♥';b.textContent=w?'♡':'♥';b.style.color=w?'':'var(--red)';toast(w?'Removed from wishlist':'Added to wishlist ♥','💝')}

function renderPayment(){
  const sub=cart.reduce((s,i)=>s+i.price*i.qty,0);
  const ship=sub>50?0:9.99;const tax=sub*.1;const tot=sub+ship+tax;
  document.getElementById('psub').textContent='$'+sub.toFixed(2);
  document.getElementById('pship').textContent=ship===0?'Free':'$'+ship.toFixed(2);
  document.getElementById('ptax').textContent='$'+tax.toFixed(2);
  document.getElementById('ptot').textContent='$'+tot.toFixed(2);
  document.getElementById('oitems').innerHTML=cart.map(i=>'<div class="oit"><img class="oimg" src="'+i.thumbnail+'" onerror="this.src=\'\'"/><div><div class="oname">'+i.title.substring(0,28)+'</div><div style="font-size:.73rem;color:var(--sub)">Qty: '+i.qty+'</div><div class="oprice">$'+(i.price*i.qty).toFixed(2)+'</div></div></div>').join('');
}

function tpf(){
  const v=document.querySelector('input[name="pay"]:checked').value;
  document.getElementById('cfields').style.display=v==='card'?'block':'none';
  document.querySelectorAll('.pmeth').forEach(e=>e.classList.remove('sel'));
  document.getElementById('pm-'+v)?.classList.add('sel');
}

function fc(inp){let v=inp.value.replace(/\D/g,'').substring(0,16);inp.value=v.replace(/(.{4})/g,'$1 ').trim()}
function fe(inp){let v=inp.value.replace(/\D/g,'');if(v.length>=2)v=v.substring(0,2)+'/'+v.substring(2);inp.value=v.substring(0,5)}

function placeOrder(){
  if(!cart.length){toast('Cart is empty!','⚠');return}
  const fn=document.getElementById('fn').value.trim();
  const em=document.getElementById('em').value.trim();
  if(!fn||!em){toast('Please fill required fields','⚠');return}
  const oid='ORD-'+Date.now().toString(36).toUpperCase();
  document.getElementById('oid').textContent=oid;
  cart=[];saveCart();
  document.getElementById('sov').classList.add('show');
}

async function initGallery(){
  if(gItems.length){rm2('all');return}
  const cm={electronics:['smartphones','laptops'],fashion:['tops','womens-bags','sunglasses'],home:['home-decoration','furniture'],beauty:['fragrances','skincare']};
  if(!allP.length)await loadP();
  gItems=allP.slice(0,60).map(p=>({src:p.images?.[0]||p.thumbnail,title:p.title,cat:Object.keys(cm).find(k=>cm[k].includes(p.category))||'home'}));
  rm2('all');
}

function rm2(filter){
  const items=filter==='all'?gItems:gItems.filter(i=>i.cat===filter);
  document.getElementById('masonry').innerHTML=items.map((item,idx)=>'<div class="mitem" onclick="olb(\''+item.src+'\')"><img src="'+item.src+'" alt="'+item.title+'" loading="lazy" style="aspect-ratio:'+(idx%3===0?'3/4':idx%3===1?'1':'4/3')+'"/><div class="movl"><span class="mlbl">'+item.title+'</span></div></div>').join('');
}

function fg(cat,btn){document.querySelectorAll('.gfb').forEach(b=>b.classList.remove('active'));btn.classList.add('active');rm2(cat)}
function olb(src){document.getElementById('lbimg').src=src;document.getElementById('lbox').classList.add('open');document.body.style.overflow='hidden'}
function clb(){document.getElementById('lbox').classList.remove('open');document.body.style.overflow=''}
function copyCode(){navigator.clipboard.writeText('ASLAN20').then(()=>toast('Code copied!','📋'))}

ucc();loadP();
</script>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":65534", nil)
}
