import { Link } from "react-router-dom";
import { useGetAllShelvesQuery } from "../api/bookstatApi";

export default function ShelvesPreviewContainer() {
  const { data: shelves, isLoading, isError } = useGetAllShelvesQuery();

  if (isLoading) {
    return <p className="text-gray-400 text-center">Ladataan hyllyjä...</p>;
  }

  if (isError) {
    return (
      <p className="text-red-500 text-center">Virhe ladattaessa hyllyjä</p>
    );
  }

  return (
    <div className="p-4">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Hyllyt</h2>
        <Link
          to="/create-shelf"
          className="px-3 py-4 bg-green-600 text-white rounded-md hover:bg-green-700"
        >
          Luo uusi hylly
        </Link>
      </div>

      {/* Shelves grid */}
      <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
        {shelves?.map((shelf) => (
          <Link
            key={shelf.id}
            to={`/shelves/${shelf.id}`}
            className="bg-gray-800 rounded-lg p-4 flex flex-col items-center justify-center hover:bg-gray-700 transition"
          >
            <span className="text-white font-medium">{shelf.name}</span>
          </Link>
        ))}
      </div>
    </div>
  );
}
