The aim of the codebase is to build some common logic to BIND DS service & Lark bot:

For now we seems to have an alpha version of it BUT we still have a long way to GO.
ISSUE for now:
1. UGLY(hard to describe the situation now but I'll give some better one later)
2. not testable
3. not maintainable as a micro service

I'll declear a clear path to complete the missing part.
stage 1: make clear interface at certain level & make some test case
stage 2: change the implementation part with no test break
stage 3: move it to our main service codebase & add maintainable feature


SUGGESTION for wangshuwen.rhett:
1. read more about context, know what's its usage & why we always use it in our service: answer what is a context tree

NOW:
stage 1, Contributor: zhangteng.dance, wangshuwen.rhett


Design for new version:
- DS client: only cares about how to deal with DS server. & export DoReq method, req struct, resp struct
- Session Component: a component of Lark Side Dialog Session, the components do not know each other, but Session know all its component
- Session: build up with Session Component, have a State Machine, knows how to trans the states.
- several Lark req&resp Translater: knows how to translate certain lark message to menu sellection or dialog content/ ds resp to lark card
- SessionManager: bind Session with Cache & provide GetSessionBySessionId (function): 
-- GetSessionBySessionId (function): the method of getting a session, it should return a session with its history state(if we have),
- Cache: 1st version will use a thread safe global key-value map, only provide Get & Set Method