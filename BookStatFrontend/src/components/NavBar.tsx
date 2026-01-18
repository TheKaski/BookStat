import { Link } from "react-router-dom";

export default function NavBar() {
  return (
    <header className="sticky top-0 bg-dark-gray-700 shadow-md px-4 py-2 pb-2 flex justify-center items-center order-last md:order-first">
      <nav className="space-x-4 flex items-center">
        <Link to="/" className="font-semibold text-white hover:text-green-600">
          Etusivu
        </Link>
        <Link
          to="/scan"
          className="font-semibold text-white hover:text-green-600"
        >
          Skannaa
        </Link>
        <Link
          to="/analytics"
          className="font-semibold text-white hover:text-green-600"
        >
          Analytiikka
        </Link>
      </nav>
    </header>
  );
}
