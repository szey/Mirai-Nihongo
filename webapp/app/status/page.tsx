export default function StatusPage() {
  return (
    <main className="page">
      <div className="container card">
        <span className="eyebrow">Project Status</span>
        <h1>Current build state</h1>
        <p className="muted">
          The repository now has a clean web-first direction, real Pre-N5 content files, a runnable Next.js scaffold,
          and the next step is to wire those content files through a Go backend API.
        </p>
        <div className="list">
          <div><strong>Direction</strong><br /><span className="muted">Web-first beginner product</span></div>
          <div><strong>Content</strong><br /><span className="muted">content/pre-n5/v1</span></div>
          <div><strong>Frontend</strong><br /><span className="muted">Next.js scaffold added</span></div>
          <div><strong>Backend</strong><br /><span className="muted">Go API scaffold is next</span></div>
        </div>
      </div>
    </main>
  );
}
